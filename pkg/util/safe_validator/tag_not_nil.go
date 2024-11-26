package safe_validator

import "fmt"

// +example:
// rkind: kChan, kFunc, kMap, kPtr, kUnsafePtr, kAny, kSlice
func (f fieldOf) validNotNil() error {
	if f.rval.IsNil() {
		return fmt.Errorf(errValInvalidFmt, f.name, f.rval.String(), "is nil")
	}
	return nil
}
