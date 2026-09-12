package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	path := strings.TrimSuffix(parsedURL.EscapedPath(), "/")
	normURL := fmt.Sprintf("%s%s", parsedURL.Host, path)
	return normURL, nil
}
