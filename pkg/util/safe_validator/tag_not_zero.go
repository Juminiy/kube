package safe_validator

import "fmt"

func (f fieldOf) validNotZero() error {
	if ptrNilErr := f.errPointerNil(notZero, ""); ptrNilErr != nil {
		return ptrNilErr
	}
	cloneF, ok := f.indirect(notZero)
	if !ok {
		return nil
	} // skip indirect value mismatch tag

	if cloneF.rval.IsZero() {
		return fmt.Errorf(errValInvalidFmt, cloneF.name, cloneF.val, "is zero")
	}
	return nil
}
