package mock

import (
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	"time"
)

type TimeFunc func() time.Time

var timeFunc = map[string]TimeFunc{}

var defaultTime = time.Now

var _timeTyp = safe_reflect.Of(time.Now()).Typ

var timeRule = rule{}
