package safe_reflectv3

type ETv struct {
	Tv
	*FieldOmit
}

func (tv Tv) ETv() *ETv {
	return &ETv{Tv: tv}
}
