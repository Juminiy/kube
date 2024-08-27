package util

import "fmt"

func SilentHandleError(handle string, err error) {
	if err != nil {
		consoleLogError(handle, err)
	}
}

func consoleLogError(detail string, err error) {
	fmt.Printf("%s: %v\n", detail, err)
}
