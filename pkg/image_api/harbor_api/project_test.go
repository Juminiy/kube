package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"testing"
)

func TestClient_ListProjects(t *testing.T) {
	listProjects, err := testHarborClient.ListProjects(true)
	util.Must(err)
	for _, projectPtr := range listProjects.Payload {
		bs, err := projectPtr.MarshalBinary()
		util.Must(err)
		t.Log(util.Bytes2StringNoCopy(bs))
	}
}

func TestClient_CreateProject(t *testing.T) {
	req := ProjectReqConfig{
		MetaDataPublic: "false",
		ProjectName:    util.StringJoin("-", "10", "200", random.LowerCaseString(10)),
		RegistryId:     0,
		StorageLimit:   50 * util.Gi,
	}
	_, err := testHarborClient.CreateProject(req)
	t.Log(testHarborClient.ErrorDetail(err))

	getProject, err := testHarborClient.GetProject(req.ProjectName)
	util.MustDetail(err)

	bs, err := getProject.Payload.MarshalBinary()
	util.MustDetail(err)
	t.Log(util.Bytes2StringNoCopy(bs))
}

func TestClient_DeleteProject(t *testing.T) {
	_, err := testHarborClient.DeleteProject("qo0ww_sa")
	util.Must(err)
}
