package util

import "testing"

func TestOSCreateAbsolutePath(t *testing.T) {
	// not absolute path
	SilentHandleError("os create absolute path", OSCreateAbsolutePath("test_dir"))

	// not exists dir
	SilentHandleError("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file"))

	// exists dir
	SilentHandleError("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file2"))

	// not exists filename
	SilentHandleError("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file3"))

	// exists filename
	SilentHandleError("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file3"))
}
