package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "", "Address to check")
	flag.Parse()

	req, err := http.NewRequest("GET", *addr, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			os.Exit(1)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	os.Exit(0)
}
