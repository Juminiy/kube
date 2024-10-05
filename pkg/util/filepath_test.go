package util

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"path/filepath"
	"testing"
)

// +passed windows darwin
func TestOSCreateAbsolutePath(t *testing.T) {
	// not absolute path
	SilentErrorf("os create absolute path", OSCreateAbsolutePath("test_dir"))

	dir0, err := internal_api.GetWorkPath("testdata")
	SilentFatal(err)

	// not exists dir
	SilentErrorf("os create absolute path", OSCreateAbsolutePath(filepath.Join(dir0, "testdir1", "test_log1.log")))

	// exists dir
	SilentErrorf("os create absolute path", OSCreateAbsolutePath(filepath.Join(dir0, "testdir0", "test_log1.log")))

	// not exists filename
	SilentErrorf("os create absolute path", OSCreateAbsolutePath(filepath.Join(dir0, "testdir0", "test_log2.log")))

	// exists filename
	SilentErrorf("os create absolute path", OSCreateAbsolutePath(filepath.Join(dir0, "testdir0", "test_log0.log")))
}
