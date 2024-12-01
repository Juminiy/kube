package safe_reflectv2

// type Value variable var is v
// type Type variable var is t
// type any variable var is i
// type reflect.Value variable var is rv
// type reflect.Type variable var is rt

func Indirect(i any) Value {
	d := Direct(i)
	d.indirect()
	return d
}

func Direct(i any) Value {
	return Value{Value: direct(i), i: i}
}
