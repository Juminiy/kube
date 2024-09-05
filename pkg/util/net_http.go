package util

import "strings"

func URLWithHTTP(url string) string {
	if strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://") {
		return url
	}
	return StringConcat("http://", url)
}
