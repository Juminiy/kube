package harbor_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/random"
	"github.com/Juminiy/kube/pkg/util/safe_json"
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
	t.Log(err)

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
	_, err := UnwrapErr(_cli.ArtifactCopyAndTag(
		ArtifactURI{
			Project:    "library",
			Repository: "jammy-release",
			Tag:        "v1.1",
		},
		ArtifactURI{
			Project:    "library",
			Repository: "jammy-env",
			Tag:        "v1.9",
			Digest:     "sha256:84b5e24c1321786da5e5ab3a56c6a06e9b181c70274fe72810a7ab0b3d064daf",
		},
	))
	util.Must(err)
	arti, err := UnwrapErr(_cli.GetArtifact(
		ArtifactURI{
			Project:    "library",
			Repository: "jammy-release",
			Tag:        "v1.1",
		},
	))
	util.Must(err)
	t.Log(safe_json.Pretty(arti))
}
