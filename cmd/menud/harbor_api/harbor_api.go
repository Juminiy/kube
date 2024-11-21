package harbor_api

import (
	"encoding"
	"fmt"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api/harbor_inst"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/zero_reflect"
)

func Menu(s ...string) {
	// image list :repo
	var (
		appOf    = s[0]
		actionOf = s[1]

		projectName    string
		repositoryName string
		pageNum        int64
		pageSize       int64
	)

	//hCli, err := harbor_api.New("", true, "", "")
	//if err != nil {
	//	fmt.Printf("new harbor repo client error: %v\n", err)
	//	return
	//}

	inputFn := func() {
		fmt.Printf("project: repository: pageNum: pageSize: ")
		_, err := fmt.Scanf("%s %s %d %d", &projectName, &repositoryName, &pageNum, &pageSize)
		if err != nil {
			fmt.Printf("harbor input project config error: %v\n", err)
			return
		}
	}

	if appOf == "project" && actionOf == "list" {
		ls, err := harbor_inst.ListProjects(false)
		if err != nil {
			fmt.Printf("harbor list proj error: %v\n", err)
			return
		}
		for _, proj := range ls.Payload {
			printBinaryMarshaller(proj)
		}
	} else if appOf == "repository" && actionOf == "list" {
		inputFn()
		ls, err := harbor_inst.ListRepositories(projectName)
		if err != nil {
			fmt.Printf("harbor list repo error: %v\n", err)
			return
		}
		for _, repo := range ls.Payload {
			printBinaryMarshaller(repo)
		}
	} else if appOf == "artifact" && actionOf == "list" {
		inputFn()
		ls, err := harbor_inst.ListArtifacts(harbor_api.ArtifactURI{
			Project:    projectName,
			Repository: repositoryName,
		})
		if err != nil {
			fmt.Printf("harbor list arti error: %v\n", err)
			return
		}
		for _, arti := range ls.Payload {
			printBinaryMarshaller(arti)
		}
	}

}

func printBinaryMarshaller(bm encoding.BinaryMarshaler) {
	bs, err := bm.MarshalBinary()
	if err != nil {
		fmt.Printf("harbor list %v error: %v\n", zero_reflect.TypeOf(bm).String(), bs)
		return
	}
	fmt.Println(util.Bytes2StringNoCopy(bs))
}

func printBinaryMarshallerList(bm []encoding.BinaryMarshaler) {
	for _, elem := range bm {
		bs, err := elem.MarshalBinary()
		if err != nil {
			fmt.Printf("harbor list %v error: %v\n", zero_reflect.TypeOf(bm).String(), bs)
			return
		}
		fmt.Println(util.Bytes2StringNoCopy(bs))
	}
}
