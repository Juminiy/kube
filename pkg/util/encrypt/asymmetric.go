package encrypt

const (
	AsymmetricRSA = iota
	AsymmetricDSA
	AsymmetricECC
)

type Asymmetric struct {
	Algo string
}
