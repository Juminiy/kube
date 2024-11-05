// Package mockv2 was generated
package mockv2

import (
	"time"
)

type TimeFunc func() time.Time

var timeFunc = map[string]TimeFunc{
	defaultKey: defaultTime,
}

var defaultTime = time.Now

var timeRule = rule{}

func (r *rule) applyTime(lval, rval string) {

}
