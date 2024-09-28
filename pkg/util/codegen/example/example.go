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

func (s *ExampleStruct) Func4(v1 int, v2 ...int) {}

func (s *ExampleStruct) Func5(v1 func(), v2 ...util.Page) {}

func (s *ExampleStruct) Func6(v1 func(), v2 string, v3 ...*string) {}

func (s *ExampleStruct) Func7(v1 []*int, v2 []int) {}

func (s *ExampleStruct) Func8(v1 *int, v2 int) {}
