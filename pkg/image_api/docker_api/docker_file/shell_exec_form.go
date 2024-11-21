package docker_file

import "io"

type ShellExecForm interface {
	Builder
	Exec(w io.StringWriter)
	Shell(w io.StringWriter)
}

type Exe struct {
	Bin   string
	Param []string
}

func (e Exe) Build(w io.StringWriter) {

}

// INSTRUCTION ["executable","param1","param2"]
func (e Exe) Exec(w io.StringWriter) {

}

// INSTRUCTION command param1 param2
func (e Exe) Shell(w io.StringWriter) {

}
