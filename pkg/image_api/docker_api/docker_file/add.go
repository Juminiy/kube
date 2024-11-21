package docker_file

// Add
// ADD [OPTIONS] <src> ... <dest>
// ADD [OPTIONS] ["<src>", ... "<dest>"]
type Add struct {
	Options []AddOption
	Src     SrcType
	Srcs    []SrcType
	Dest    DestType
}

type AddOption string

const (
	KeepGitDir AddOption = "--keep-git-dir" //	1.1
	Checksum   AddOption = "--checksum"     //	1.6
	Chown      AddOption = "--chown"        //
	Chmod      AddOption = "--chmod"        //1.2
	Link       AddOption = "--link"         //1.4
	Exclude    AddOption = "--exclude"      //1.7-labs
)

type SrcType string

type DestType string
