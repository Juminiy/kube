package example

import "github.com/Juminiy/kube/pkg/util"

type ExampleStruct struct {
	Field1 int
	Field2 string
}

func (s *ExampleStruct) Func1(int, string) {

}

func (s *ExampleStruct) Func2(page *util.Page) util.Page {
	return *page
}

func (s *ExampleStruct) Func3(pagePtr *util.Page, pageVal util.Page) (*util.Page, util.Page) {
	return &pageVal, *pagePtr
}
