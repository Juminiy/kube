package harbor_api

import (
	"encoding"
	"fmt"
	"kube/util"
	"reflect"
)

func Menu(s ...string) {
	// image list :repo
	var (
		appOf    = s[0]
		actionOf = s[1]
	)

	repoClient, err := NewHarborCli()
	if err != nil {
		fmt.Printf("new harbor repo client error: %v\n", err)
		return
	}

	if appOf == "project" && actionOf == "list" {
		ls, err := repoClient.ListProjects()
		if err != nil {
			fmt.Printf("harbor list project error: %v\n", err)
			return
		}
		for _, proj := range ls.Payload {
			printBinaryMarshaler(proj)
		}
	} else if appOf == "image" && actionOf == "list" {
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

		ls, err := repoClient.
			WithProjectName(projectName).
			WithPageConfig(util.NewPageConfig(pageNum, pageSize)).
			ListImages()
		if err != nil {
			fmt.Printf("harbor list image error: %v\n", err)
			return
		}
		for _, image := range ls.Payload {
			printBinaryMarshaler(image)
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
