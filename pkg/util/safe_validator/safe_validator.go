package safe_validator

// Validate
// compatible with fiber.StructValidator
func (cfg *Config) Validate(out any) error {
	return cfg.StructE(out)
}
