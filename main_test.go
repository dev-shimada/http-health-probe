package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestParseAddr(t *testing.T) {
	testCases := []struct {
		name     string
		addr     string
		useTLS   bool
		expected string
	}{
		{
			name:     "with http prefix",
			addr:     "http://example.com",
			useTLS:   false,
			expected: "http://example.com",
		},
		{
			name:     "with https prefix",
			addr:     "https://example.com",
			useTLS:   true,
			expected: "https://example.com",
		},
		{
			name:     "without prefix, no tls",
			addr:     "example.com",
			useTLS:   false,
			expected: "http://example.com",
		},
		{
			name:     "without prefix, with tls",
			addr:     "example.com",
			useTLS:   true,
			expected: "https://example.com",
		},
		{
			name:     "with port, no tls",
			addr:     ":8080",
			useTLS:   false,
			expected: "http://localhost:8080",
		},
		{
			name:     "with port, with tls",
			addr:     ":8080",
			useTLS:   true,
			expected: "https://localhost:8080",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseAddr(tc.addr, tc.useTLS)
			if actual != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		cfg := config{
			addr:           server.URL,
			method:         "GET",
			timeout:        1 * time.Second,
			expectedStatus: http.StatusOK,
		}

		exitCode := run(cfg)

		if exitCode != ExitOK {
			t.Errorf("expected exit code %d, but got %d", ExitOK, exitCode)
		}
	})

	t.Run("unexpected status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		cfg := config{
			addr:           server.URL,
			method:         "GET",
			timeout:        1 * time.Second,
			expectedStatus: http.StatusOK,
		}

		exitCode := run(cfg)

		if exitCode != ExitBadStatus {
			t.Errorf("expected exit code %d, but got %d", ExitBadStatus, exitCode)
		}
	})

	t.Run("connection failed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		server.Close() // Close server to simulate connection failure

		cfg := config{
			addr:           server.URL,
			method:         "GET",
			timeout:        1 * time.Second,
			expectedStatus: http.StatusOK,
		}

		exitCode := run(cfg)

		if exitCode != ExitConnectionFailed {
			t.Errorf("expected exit code %d, but got %d", ExitConnectionFailed, exitCode)
		}
	})

	t.Run("invalid address", func(t *testing.T) {
		cfg := config{
			addr: "", // Invalid address
		}

		exitCode := run(cfg)

		if exitCode != ExitInvalidArgs {
			t.Errorf("expected exit code %d, but got %d", ExitInvalidArgs, exitCode)
		}
	})
}
