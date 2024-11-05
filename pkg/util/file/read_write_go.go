package file

import (
	"github.com/Juminiy/kube/pkg/util"
	"strings"
)

type (
	Go struct {
		GoSrc
		GoDst
	}

	GoSrc struct {
		// package xxx
		Package string

		// import ""
		// import ()
		Import []string

		// Text Contain Comment
		Text    []string
		Comment []string
	}

	GoDst struct {
		Comments map[string][]string
	}
)

func NewGoFile() *Go {
	return &Go{
		GoSrc: GoSrc{
			Import:  make([]string, 0, 8),
			Text:    make([]string, 0, 256),
			Comment: make([]string, 0, 64),
		},
		GoDst: GoDst{
			Comments: make(map[string][]string, 8), // write comment kinds
		},
	}
}

func ReadGo(filePath string) *Go {
	rd := NewReader(filePath)
	var goFile = NewGoFile()
	var lineStr string
	for rd.NextLine(&lineStr) {
		if strings.HasPrefix(lineStr, "//") {
			goFile.Comment = append(goFile.Comment, lineStr)
		}
		switch {
		case strings.HasPrefix(lineStr, "package"):
			goFile.Package = strDelete(lineStr, "package ")

		case strings.HasPrefix(lineStr, "import"):
			// import ""
			// import _ ""
			if !strings.Contains(lineStr, "(") {
				goFile.Import = append(goFile.Import, strDelete(lineStr, "import "))
			} else {
				// import (
				for rd.NextLine(&lineStr) {
					// )
					if strings.HasPrefix(lineStr, ")") {
						break
					} else {
						goFile.Import = append(goFile.Import, lineStr)
					}
				}
			}

		default:
			goFile.Text = append(goFile.Text, lineStr)
		}

	}
	return goFile
}

var strDelete = util.StringDelete

func (f *Go) TrimComment() *Go {
	return f
}

func (f *Go) PackageOf(packageOf string) *Go {
	f.Package = packageOf
	f.Comments[CommentBeforePackage] = append(
		f.Comments[CommentBeforePackage], "//Package "+packageOf+" was"+" generated")
	return f
}

func (f *Go) WriteTo(dstFilePath string) {
	wr := NewWriter(dstFilePath).SkipSpace()

	// package
	if beforePackage, ok := util.MapElemOk(f.Comments, CommentBeforePackage); ok {
		wr.WordsSep("\n", beforePackage...)
	}
	wr.Words("package ", f.Package, "\n\n")

	// import
	if beforePackage, ok := util.MapElemOk(f.Comments, CommentBeforeImport); ok {
		wr.WordsSep("\n", beforePackage...)
	}
	wr.Words("import", "(", "\n")
	wr.Words(f.Import...)
	wr.Word(")", "\n")

	// text
	wr.Words(f.Text...)

	wr.Done()
}

const (
	CommentBeforePackage = "before:package"
	CommentBeforeImport  = "before:import"
)
