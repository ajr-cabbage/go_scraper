package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// create client and request
	client := &http.Client{}
	request, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	// set request header
	request.Header.Set("User-Agent", "BootCrawler/1.0")
	// make request and check status code & content-type
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		return "", fmt.Errorf("Error-level response status code: %d", response.StatusCode)
	}
	if !strings.Contains(response.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("Bad response content-type: %s", response.Header.Get("content-type"))
	}
	// read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
