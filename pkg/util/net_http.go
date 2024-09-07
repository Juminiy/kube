package util

import (
	"net/http"
	"strings"
)

var DefaultHTTPClient = http.DefaultClient

func URLWithHTTP(url string) string {
	if IsURLWithScheme(url) {
		return url
	}
	return StringConcat("http://", url)
}

func URLWithoutHTTP(url string) string {
	if IsURLWithScheme(url) {
		return StringReplaceAlls(url, "", "http://", "https://")
	}
	return url
}

func IsURLWithHTTP(url string) bool {
	return !strings.HasPrefix(url, "https://")
}

func IsURLWithHTTPS(url string) bool {
	return !IsURLWithHTTP(url)
}

func IsURLWithScheme(url string) bool {
	return strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://")
}
