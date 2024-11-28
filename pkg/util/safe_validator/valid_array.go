package safe_validator

func (cfg *Config) Array(v any, tag ...string) bool {
	return true
}

func (cfg *Config) ArrayE(v any, tag ...string) error {
	return nil
}
