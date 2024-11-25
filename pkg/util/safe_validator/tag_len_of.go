package safe_validator

import "fmt"

// +param: tagv
// +example:
// rkind: kArr, kChan, kMap, kSlice, kString
// tagv			| byRange
// len:10		| 10~10
// len:~20		| 0~20
// len:30~ 		| 30~math.MaxInt
// len:20~60 	| 20~60
// len:-1~100	| 0~100
// len:-5~-10	| error
// len:11~2		| error
func (f fieldOf) validLen(tagv string) error {
	return nil
}

func lenFormatErr(tagv string) error {
	return fmt.Errorf("len format error: (%s)", tagv)
}

func lenValidErr(lenRange string, v any) error {
	return fmt.Errorf("len valid error: %v not in range (%s)", v, lenRange)
}

func parseRange() {

}
