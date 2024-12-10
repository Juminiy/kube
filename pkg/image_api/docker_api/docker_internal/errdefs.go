package docker_internal

import (
	_ "github.com/docker/docker/errdefs"
	"strings"
)

// Modify of github.com/docker/docker/errdefs

type ErrParser struct{}

var _errParser ErrParser

func GetErrParser() ErrParser {
	return _errParser
}

// +example:
// "Error response from daemon: unknown: artifact library/hello:v1.0 not found"
func (ErrParser) ArtifactNotFound(err error) bool {
	if err != nil {
		words := strings.Split(err.Error(), " ")
		indexArti, indexNot, indexFound := -1, -1, -1
		for i := range words {
			switch words[i] {
			case "artifact":
				indexArti = i
			case "not":
				indexNot = i
			case "found":
				indexFound = i
			}
		}
		return indexArti > 0 && indexNot > 0 && indexFound > 0 && indexFound == indexNot+1
	}
	return false
}
