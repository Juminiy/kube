//go:build windows

package docker_api

const (
	testTarGzPath                = "testdata\\tar_gz\\hello-world.tar.gz"
	testTarGZPathExportedByLinux = "testdata\\tar_gz\\hello-world-latest.tar.gz"
	testTarBuildPath             = "testdata\\tar_gz\\webapp.tar"
	testTarBuildTimeout          = "testdata\\tar_gz\\timeout.tar"
)