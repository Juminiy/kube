package safe_validator

func (cfg *Config) Map(v any) bool {
	return true
}

func (cfg *Config) MapE(v any) error {
	return nil
}
