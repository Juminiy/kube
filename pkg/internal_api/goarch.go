package internal_api

// referred from: internal/goarch

// copy from: internal/goarch/goarch.go
type ArchFamilyType int

const (
	AMD64 ArchFamilyType = iota
	ARM
	ARM64
	I386
	LOONG64
	MIPS
	MIPS64
	PPC64
	RISCV64
	S390X
	WASM
)

const (
	_386        = `386`
	Amd64       = `amd64`
	Amd64p32    = ``
	Arm         = `arm`
	Armbe       = ``
	Arm64       = `arm64`
	Arm64be     = `arm64be`
	Loong64     = `loong64`
	Mips        = `mips`
	Mipsle      = `mipsle`
	Mips64      = `mips64`
	Mips64le    = `mips64le`
	Mips64p32   = `mips64p32`
	Mips64p32le = `mips64p32le`
	Ppc         = `ppc`
	Ppc64       = `ppc64`
	Ppc64le     = `ppc64le`
	Riscv       = `riscv`
	Riscv64     = `riscv64`
	S390        = `s390`
	S390x       = `s390x`
	Sparc       = `sparc`
	Sparc64     = `sparc64`
	Wasm        = `wasm`
)
