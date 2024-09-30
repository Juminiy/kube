package internal_api

import "testing"

// +passed +windows +darwin
func TestGetWorkPath(t *testing.T) {
	t.Log(GetWorkPath("testdata", "test_log"))
}

// +passed +windows +darwin
func TestOpenFileWithCreateIfNotExist(t *testing.T) {
	filePtr, err := OpenFileWithCreateIfNotExist(filePathJoin("testdata", "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filePtr.Fd())
}
