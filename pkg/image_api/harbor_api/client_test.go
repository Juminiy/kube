package harbor_api

var (
	testHarborClient, testHarborClientErr = New(
		"http://harbor.local:18111",
		true,
		"admin",
		"bupt.harbor@666",
	)
)
