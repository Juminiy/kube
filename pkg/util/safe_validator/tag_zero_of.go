package safe_validator

import "fmt"

func (f fieldOf) validNotZero() error {
	if f.rval.IsZero() {
		return fmt.Errorf(errValInvalidFmt, f.name, f.val, "is zero")
	}
	return nil
}

func (f fieldOf) validIsZero() error {
	if !f.rval.IsZero() {
		return fmt.Errorf(errValInvalidFmt, f.name, f.val, "is not zero")
	}
	return nil
}
