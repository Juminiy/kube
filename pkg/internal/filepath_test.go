package internal

import "testing"

// +passed +windows
func TestGetWorkPath(t *testing.T) {
	t.Log(GetWorkPath("pkg", "log_api", "stdlog", "testdata", "test_log"))
}

// +passed +windows
func TestOpenFileWithCreateIfNotExist(t *testing.T) {
	filePtr, err := OpenFileWithCreateIfNotExist(filePathJoin("testdata", "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filePtr.Fd())
}
