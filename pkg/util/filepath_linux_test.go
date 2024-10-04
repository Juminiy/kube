//go:build linux

package util

func TestOSCreateAbsolutePath(t *testing.T) {
	// not absolute path
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("test_dir"))

	// not exists dir
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file"))

	// exists dir
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file2"))

	// not exists filename
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file3"))

	// exists filename
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("/home/wz/test_dir/test_file3"))
}
