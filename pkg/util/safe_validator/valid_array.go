package safe_validator

import "errors"

func (cfg *Config) Array(v any) bool {
	ok, _ := cfg.arrayOkE(v)
	return ok
}

func (cfg *Config) ArrayE(v any) error {
	_, err := cfg.arrayOkE(v)
	return err
}

func (cfg *Config) arrayOkE(v any) (ok bool, err error) {
	tv := indir(v)
	if tv.Val.Kind() != kArr {
		return true, errValNotArray
	}

	return cfg.arrayLikeOkE(tv)
}

var errValNotArray = errors.New("value type is not array")
var errArrayElemNotStruct = errors.New("array elem type is not struct")
