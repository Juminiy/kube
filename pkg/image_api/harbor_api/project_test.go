package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"testing"
)

func TestClient_ListProjects(t *testing.T) {
	listProjects, err := _cli.ListProjects(true)
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
		StorageLimit:   50 * util.Gi,
	}
	_, err := _cli.CreateProject(req)
	t.Log(_cli.ErrorDetail(err))

	getProject, err := _cli.GetProject(req.ProjectName)
	util.MustDetail(err)

	bs, err := getProject.Payload.MarshalBinary()
	util.MustDetail(err)
	t.Log(util.Bytes2StringNoCopy(bs))
}

func TestClient_DeleteProject(t *testing.T) {
	_, err := _cli.DeleteProject("qo0ww_sa")
	util.Must(err)
}

func TestClient_CopyArtifact(t *testing.T) {
	_, err := UnwrapErr(_cli.CopyArtifact(
		ArtifactURI{
			Project:    "library",
			Repository: "ubuntu",
			Tag:        "jammy",
		},
		ArtifactURI{
			Project:    "k8e",
			Repository: "ubuntu",
			Tag:        "jammy",
		},
	))
	t.Log(_cli.ErrorDetail(err))
}
