//go:build unix

package util

import "testing"

// Deprecated: github.com/Juminiy/kube/pkg/util/filepath_unix.go do not maintain anymore
// use: github.com/Juminiy/kube/pkg/internal/filepath_unix.go instead

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
