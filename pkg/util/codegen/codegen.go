package codegen

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"os"
	"reflect"
	"strings"
)

type Manifest struct {
	// +optional
	UnExportGlobalVarName string

	// InstanceOf only allow pointer type (Addressable type: any, func, pointer, map, slice)
	// +required
	InstanceOf any

	// +required
	DstFilePath string

	// DstPkg must not be same with InstanceOf pkg
	// DstPkg must same with DstFilePath pkg (or suffix word)
	// +required
	DstPkg string

	GenPackage bool
	GenImport  bool
	GenVar     bool

	file *os.File
	reflectType
}

func (g *Manifest) GetDstPkg() string {
	parts := strings.Split(g.DstFilePath, "/")
	if len(parts) <= 1 {
		return g.DstPkg
	}
	g.DstPkg = parts[len(parts)-2]
	return g.DstPkg
}

func (g *Manifest) Write(s string) *Manifest {
	_, err := g.file.WriteString(s)
	util.SilentPanicError(err)
	_, err = g.file.WriteString(" ")
	util.SilentPanicError(err)
	return g
}

func (g *Manifest) Writes(s ...string) *Manifest {
	for _, elem := range s {
		g.Write(elem)
	}
	return g
}

func (g *Manifest) NextLine() *Manifest {
	_, err := g.file.WriteString("\n")
	util.SilentPanicError(err)
	return g
}

func (g *Manifest) Do() {
	if g.InstanceOf == nil {
		stdlog.Error("InstanceOf is nil, do nothing, exit 1")
		os.Exit(1)
	}
	g.reflectType.inst = g.InstanceOf
	g.reflectType.parse()
	g.GetDstPkg()

	var errorOf error
	g.file, errorOf = util.OSOpenFile(g.DstFilePath)
	util.SilentPanicError(errorOf)

	g.Writes("//generate by codegen, do not modify it")

	g.genPackage()

	g.genImport()

	g.genVar()

	g.genFunc()
}

func (g *Manifest) genPackage() {
	if g.GenPackage {
		g.Writes("package", g.DstPkg).NextLine()
	}
}

func (g *Manifest) genImport() {
	if g.GenImport {
		g.Writes("import", "(").NextLine()
		// need to fix, maybe gofmt cmd auto fix
		g.Writes(")").NextLine()
	}
}

func (g *Manifest) genVar() {
	if g.GenVar {
		g.Writes("var", "(").NextLine()
		g.Writes(g.UnExportGlobalVarName, g.pkgTypePtrName).NextLine()
		g.Writes(")").NextLine()
		g.NextLine()
	}
}

func (g *Manifest) genFunc() {
	// for each func
	for i := range g.typ.NumMethod() {
		method := g.typ.Method(i)
		g.Writes("func", method.Name, "(")

		// make a func instance and reflect func instance
		funcInstanceTypeOf := reflect.New(method.Type).Type().Elem()
		hasParam := funcInstanceTypeOf.NumIn() > 0
		hasRetVal := funcInstanceTypeOf.NumOut() > 0

		paramList := make([]string, 0, funcInstanceTypeOf.NumIn())

		// func parameter in
		for j := range funcInstanceTypeOf.NumIn() {
			paramIn := funcInstanceTypeOf.In(j)
			if j == 0 {
				continue
			}
			paramInTyp := reflectType{typ: paramIn}
			paramInTyp.parse()
			g.Writes(paramInTyp.varName, paramInTyp.fullTypeName)
			paramList = append(paramList, paramInTyp.varName)
			if j < funcInstanceTypeOf.NumIn()-1 {
				g.Writes(",")
			}
		}

		if !hasRetVal {
			g.Writes(")", "{").NextLine()
		} else {
			g.Writes(")", "(")
			// func return value out
			for j := range funcInstanceTypeOf.NumOut() {
				retValOut := funcInstanceTypeOf.Out(j)
				retValOutTyp := reflectType{typ: retValOut}
				retValOutTyp.parse()
				g.Writes(retValOutTyp.fullTypeName)
				if j < funcInstanceTypeOf.NumOut()-1 {
					g.Writes(",")
				}
			}
			g.Writes(")", "{").NextLine()
		}

		// func body
		if hasRetVal {
			g.Writes("return")
		}
		g.Writes(g.UnExportGlobalVarName + "." + method.Name + "(")
		if hasParam {
			for k, paramRename := range paramList {
				g.Writes(paramRename)
				if k < len(paramList)-1 {
					g.Writes(",")
				}
			}
		}
		g.Writes(")").NextLine()

		g.Writes("}").NextLine()

		g.NextLine()
	}
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
