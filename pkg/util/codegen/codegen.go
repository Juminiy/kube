package codegen

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Manifest
// for all `func` of a `struct` codegen to `pkg` level func with a `pkg` level global var
type Manifest struct {
	// codegen destination file path
	// +required
	DstFilePath string

	// DstPkg must not be same with InstanceOf pkg
	// DstPkg must same with DstFilePath pkg (or suffix word)
	// +optional
	DstPkg string

	// InstanceOf only allow pointer type (Addressable type: any, func, pointer, map, slice)
	// +required
	InstanceOf any

	// +optional
	UnExportGlobalVarName string

	// +optional
	GenImport bool

	// +optional
	GenVar bool

	file     *os.File
	fileName string
	reflectType
}

func (g *Manifest) Do() {
	if len(g.DstFilePath) == 0 {
		util.SilentPanicError(fmt.Errorf("DstFilePath is nil, do nothing, exit 1"))
	}
	errorOf := util.OSRemoveFile(g.DstFilePath)
	util.SilentPanicError(errorOf)
	g.file, errorOf = util.OSOpenFileWithCreate(g.DstFilePath)
	util.SilentPanicError(errorOf)
	_, g.fileName = g.getDstPkgAndFileName()

	if g.InstanceOf == nil {
		util.SilentPanicError(fmt.Errorf("InstanceOf is nil, do nothing, exit 1"))
	}
	g.reflectType.inst = g.InstanceOf
	g.reflectType.parse()

	if len(g.UnExportGlobalVarName) == 0 {
		g.UnExportGlobalVarName = "_global" + g.typeName
	}

	// codegen comment header
	g.writes("//", "Package", g.DstPkg+"/"+g.fileName, "was generated by codegen, do not modify it").nextLine()

	g.genPackage()

	g.genImport()

	g.genVar()

	g.genFunc()
}

func (g *Manifest) genPackage() {
	g.writes("package", g.DstPkg).nextLine()
}

func (g *Manifest) genImport() {
	if g.GenImport {
		g.writes("import", "(").nextLine()
		// need to fix, maybe gofmt cmd auto fix
		g.writes(")").nextLine()
	}
}

func (g *Manifest) genVar() {
	if g.GenVar {
		g.writes("var", "(").nextLine()
		g.writes(g.UnExportGlobalVarName, g.pkgTypePtrName).nextLine()
		g.writes(")").nextLine()
		g.nextLine()
	}
}

func (g *Manifest) genFunc() {
	// for each func
	for i := range g.typ.NumMethod() {
		method := g.typ.Method(i)
		g.writes("func", method.Name, "(")

		// make a func instance and reflect func instance
		funcInstanceTypeOf := reflect.New(method.Type).Type().Elem()
		hasParam := funcInstanceTypeOf.NumIn() > 0
		hasRetVal := funcInstanceTypeOf.NumOut() > 0

		paramList := make([]string, 0, funcInstanceTypeOf.NumIn())

		fullTypeNames := make(map[string]int, funcInstanceTypeOf.NumIn())
		// func parameter in
		for j := range funcInstanceTypeOf.NumIn() {
			paramIn := funcInstanceTypeOf.In(j)
			if j == 0 {
				continue
			}
			paramInTyp := reflectType{typ: paramIn}
			paramInTyp.parse()

			paramInFullTypeName := paramInTyp.fullTypeName
			paramInVarName := paramInTyp.varName
			if inTypTh, inTypExists := fullTypeNames[paramInFullTypeName]; inTypExists {
				paramInVarName += strconv.Itoa(inTypTh + 1)
				fullTypeNames[paramInFullTypeName] = inTypTh + 1
			} else {
				fullTypeNames[paramInFullTypeName] = 1
			}
			g.writes(paramInVarName, paramInFullTypeName)

			paramList = append(paramList, paramInVarName)
			if j < funcInstanceTypeOf.NumIn()-1 {
				g.writes(",")
			}
		}

		if !hasRetVal {
			g.writes(")", "{").nextLine()
		} else {
			g.writes(")", "(")
			// func return value out
			for j := range funcInstanceTypeOf.NumOut() {
				retValOut := funcInstanceTypeOf.Out(j)
				retValOutTyp := reflectType{typ: retValOut}
				retValOutTyp.parse()
				g.writes(retValOutTyp.fullTypeName)
				if j < funcInstanceTypeOf.NumOut()-1 {
					g.writes(",")
				}
			}
			g.writes(")", "{").nextLine()
		}

		// func body
		if hasRetVal {
			g.writes("return")
		}
		g.writes(g.UnExportGlobalVarName + "." + method.Name + "(")
		if hasParam {
			for k, paramRename := range paramList {
				g.writes(paramRename)
				if k < len(paramList)-1 {
					g.writes(",")
				}
			}
		}
		g.writes(")").nextLine()

		g.writes("}").nextLine()

		g.nextLine()
	}
}

func (g *Manifest) getDstPkgAndFileName() (string, string) {
	parts := strings.Split(g.DstFilePath, "/")
	if len(parts) <= 1 {
		return g.DstPkg, parts[0]
	}
	g.DstPkg = parts[len(parts)-2]
	return g.DstPkg, parts[len(parts)-1]
}

func (g *Manifest) write(s string) *Manifest {
	_, err := g.file.WriteString(s)
	util.SilentPanicError(err)
	_, err = g.file.WriteString(" ")
	util.SilentPanicError(err)
	return g
}

func (g *Manifest) writes(s ...string) *Manifest {
	for _, elem := range s {
		g.write(elem)
	}
	return g
}

func (g *Manifest) nextLine() *Manifest {
	_, err := g.file.WriteString("\n")
	util.SilentPanicError(err)
	return g
}

type reflectType struct {
	inst any
	typ  reflect.Type

	comparePkgName string
	pkgName        string
	shortPkg       string

	fullTypeName   string
	pkgTypeName    string
	pkgTypePtrName string
	typeName       string
	typePtrName    string

	varName string
}

func (t *reflectType) parse() {
	if t.inst != nil {
		t.typ = reflect.TypeOf(t.inst)
	}

	t.pkgName, t.typeName = splitPkgType(t.typ.String())
	t.shortPkg = getLastSlash(t.pkgName)

	// pkg1/pkg2/pkg3/example.ExampleStruct
	// *pkg1/pkg2/pkg3/example.ExampleStruct
	// pkg1/pkg2/pkg3/*example.ExampleStruct

	t.fullTypeName = t.typ.String()
	t.pkgTypeName = getLastSlash(strings.ReplaceAll(t.fullTypeName, "*", ""))
	t.pkgTypePtrName = "*" + t.pkgTypeName
	t.typePtrName = "*" + t.typeName

	if strings.ToLower(string(t.typeName[0])) == string(t.typeName[0]) {
		t.varName = "_" + t.typeName
	} else {
		t.varName = strings.ToLower(string(t.typeName[0])) + t.typeName[1:]
	}

}

func getLastSlash(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func splitPkgType(s string) (string, string) {
	s = strings.ReplaceAll(s, "*", "")
	parts := strings.Split(s, ".")
	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return "", parts[0]
	}
	return parts[len(parts)-2], parts[len(parts)-1]
}

type reflectValue struct {
	inst any
	val  reflect.Value
}

func (v *reflectValue) parse() {

}