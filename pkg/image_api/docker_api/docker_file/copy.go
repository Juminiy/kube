package docker_file

// Copy
// COPY [OPTIONS] <src> ... <dest>
// COPY [OPTIONS] ["<src>", ... "<dest>"]
type Copy struct {
	Options []CopyOption
	Src     SrcType
	Srcs    []SrcType
	Dest    DestType
}

type CopyOption string

const (
	CopyOptFrom    CopyOption = "--from"
	CopyOptChown   CopyOption = "--chown"
	CopyOptChmod   CopyOption = "--chmod"   //	1.2
	CopyOptLink    CopyOption = "--link"    // 1.4
	CopyOptParents CopyOption = "--parents" //1.7-labs
	CopyOptExclude CopyOption = "--exclude" // 1.7-labs
)
