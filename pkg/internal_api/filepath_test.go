package internal_api

import "testing"

// +passed +windows +darwin
func TestGetWorkPath(t *testing.T) {
	t.Log(GetWorkPath("testdata", "test_log"))
}

// +passed +darwin
func TestAppendCreateFile(t *testing.T) {
	wd, err := GetWorkPath("testdata", "test_log", "a.txt")
	silentFatal(t, err)
	filePtr, err := AppendCreateFile(wd)
	silentFatal(t, err)
	defer filePtr.Close()
	_, err = filePtr.WriteString("\nrrr")
	silentFatal(t, err)
}

// +passed +darwin
func TestOverwriteCreateFile(t *testing.T) {
	wd, err := GetWorkPath("testdata", "test_log", "b.txt")
	silentFatal(t, err)
	filePtr, err := OverwriteCreateFile(wd)
	silentFatal(t, err)
	defer filePtr.Close()
	_, err = filePtr.WriteString("\nrrr")
	silentFatal(t, err)
}

// +passed +darwin
func TestFileExist(t *testing.T) {
	// exist
	wd, err := GetWorkPath("testdata", "test_log", "b.txt")
	silentFatal(t, err)
	t.Log(FileExist(wd))    // true
	t.Log(FileNotExist(wd)) // false

	// not exist
	wd, err = GetWorkPath("testdata", "test_log", "c.txt")
	silentFatal(t, err)
	t.Log(FileExist(wd))    // false
	t.Log(FileNotExist(wd)) // true
}

// +passed +darwin
func TestDirExist(t *testing.T) {
	// exist
	wd, err := GetWorkPath("testdata", "test_log")
	silentFatal(t, err)
	t.Log(DirExist(wd))
	t.Log(DirNotExist(wd))

	// not exist
	wd, err = GetWorkPath("testdata", "test_log_2")
	silentFatal(t, err)
	t.Log(DirExist(wd))
	t.Log(DirNotExist(wd))
}

func silentFatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
