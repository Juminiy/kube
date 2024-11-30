package safe_validator

func (cfg *Config) Any(v any) bool {

	return true
}

func (cfg *Config) AnyE(v any) error {
	return nil
}
