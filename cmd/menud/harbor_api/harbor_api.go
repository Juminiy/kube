package harbor_api

import (
	"encoding"
	"fmt"
	"kube/pkg/harbor_api"
	"kube/pkg/util"
	"reflect"
)

func Menu(s ...string) {
	// image list :repo
	var (
		appOf    = s[0]
		actionOf = s[1]
	)

	hCli, err := harbor_api.NewHarborCli()
	if err != nil {
		fmt.Printf("new harbor repo client error: %v\n", err)
		return
	}

	if appOf == "project" && actionOf == "list" {
		ls, err := hCli.ListProjects()
		if err != nil {
			fmt.Printf("harbor list project error: %v\n", err)
			return
		}
		for _, proj := range ls.Payload {
			printBinaryMarshaler(proj)
		}
	} else if appOf == "repository" && actionOf == "list" {
		var (
			projectName string
			pageNum     int64
			pageSize    int64
		)
		fmt.Printf("projectName: pageNum: pageSize: ")
		_, err := fmt.Scanf("%s %d %d", &projectName, &pageNum, &pageSize)
		if err != nil {
			fmt.Printf("harbor input project config error: %v\n", err)
			return
		}

		ls, err := hCli.
			WithPageConfig(util.NewPageConfig(pageNum, pageSize)).
			ListRepositories(projectName)
		if err != nil {
			fmt.Printf("harbor list image error: %v\n", err)
			return
		}
		for _, repo := range ls.Payload {
			printBinaryMarshaler(repo)
		}
	}

}

func printBinaryMarshaler(bm encoding.BinaryMarshaler) {
	bs, err := bm.MarshalBinary()
	if err != nil {
		fmt.Printf("harbor list %v error: %v\n", reflect.TypeOf(bm).String(), bs)
		return
	}
	fmt.Println(string(bs))
}

func printBinaryMarshalerList(bm []encoding.BinaryMarshaler) {
	for _, elem := range bm {
		bs, err := elem.MarshalBinary()
		if err != nil {
			fmt.Printf("harbor list %v error: %v\n", reflect.TypeOf(bm).String(), bs)
			return
		}
		fmt.Println(string(bs))
	}
}
