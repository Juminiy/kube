package safe_validator

func (cfg *Config) Slice(v any) bool {
	return true
}

func (cfg *Config) SliceE(v any) error {
	return nil
}
