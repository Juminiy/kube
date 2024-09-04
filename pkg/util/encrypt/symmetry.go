package encrypt

const (
	SymmetryAES = iota
	SymmetryDES
	Symmetry3DES
	SymmetryBlowfish
)

type Symmetry struct {
	Algo string
}
