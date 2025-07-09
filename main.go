package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	// ExitOK for exit code
	ExitOK int = 0
	// ExitInvalidArgs for exit code
	ExitInvalidArgs = 1
	// ExitConnectionFailed for exit code
	ExitConnectionFailed = 2
	// ExitBadStatus for exit code
	ExitBadStatus = 4
)

type config struct {
	addr           string
	method         string
	timeout        time.Duration
	expectedStatus int
	useTLS         bool
	insecure       bool
	userAgent      string
}

func main() {
	addr := flag.String("addr", "", "Address to check")
	method := flag.String("method", "GET", "HTTP method")
	timeout := flag.Duration("timeout", 1*time.Second, "Request timeout")
	expectedStatus := flag.Int("expected-status", http.StatusOK, "Expected HTTP status code")
	useTLS := flag.Bool("tls", false, "Use TLS")
	insecure := flag.Bool("insecure", false, "Skip TLS certificate verification")
	userAgent := flag.String("user-agent", "http-health-probe", "User-Agent header")

	flag.Parse()

	cfg := config{
		addr:           *addr,
		method:         *method,
		timeout:        *timeout,
		expectedStatus: *expectedStatus,
		useTLS:         *useTLS,
		insecure:       *insecure,
		userAgent:      *userAgent,
	}

	os.Exit(run(cfg))
}

func run(cfg config) int {
	if cfg.addr == "" {
		fmt.Fprintln(os.Stderr, "Error: -addr is required")
		return ExitInvalidArgs
	}

	url := parseAddr(cfg.addr, cfg.useTLS)

	req, err := http.NewRequest(cfg.method, url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating request:", err)
		return ExitConnectionFailed
	}
	req.Header.Set("User-Agent", cfg.userAgent)

	client := newClient(cfg.insecure, cfg.timeout)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error making request:", err)
		return ExitConnectionFailed
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Error closing response body:", err)
		}
	}()

	if resp.StatusCode != cfg.expectedStatus {
		fmt.Fprintf(os.Stderr, "Error: Unexpected status code: got %d, want %d\n", resp.StatusCode, cfg.expectedStatus)
		return ExitBadStatus
	}

	fmt.Println("OK")
	return ExitOK
}

func newClient(insecure bool, timeout time.Duration) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
}


func parseAddr(addr string, useTLS bool) string {
	if strings.HasPrefix(addr, ":") {
		addr = "localhost" + addr
	}
	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		if useTLS {
			addr = "https://" + addr
		} else {
			addr = "http://" + addr
		}
	}
	return addr
}
