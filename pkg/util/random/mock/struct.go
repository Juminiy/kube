package mock

func Struct(v any) {
	indir(v).StructParseTag(mockTag)
}
