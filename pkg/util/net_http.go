package util

import (
	"net/http"
	"strings"
)

var DefaultHTTPClient = http.DefaultClient

func URLWithHTTP(url string) string {
	if strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://") {
		return url
	}
	return StringConcat("http://", url)
}
