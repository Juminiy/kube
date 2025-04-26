package util

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	unit "github.com/docker/go-units"
	"github.com/prometheus/prometheus/model/timestamp"
	"time"
)

const (
	TimeLocationAsiaShanghai = "Asia/Shanghai"
)

// for human read
const (
	DurationMinute      = time.Second * 60
	DurationHour        = DurationMinute * 60
	DurationHalfDay     = DurationHour * 12
	DurationDay         = DurationHalfDay * 2
	DurationWeek        = DurationDay * 7
	DurationMonth       = DurationDay * 30
	DurationYear        = DurationDay * 365
	DurationEra         = DurationYear * 10
	DurationHalfCentury = DurationEra * 5
	DurationCentury     = DurationHalfCentury * 2
)

// CST2Timestamp
// Convert (CST: UTC+8) to (Unix timestamp)
func CST2Timestamp(cst string) string {
	cstTm, err := time.ParseInLocation(time.DateTime, cst, cstLocation())
	if err != nil {
		stdlog.ErrorF("convert cst: %s to timestamp error: %s", cst, err.Error())
		return ""
	}
	return I64toa(cstTm.Unix())
}

// Timestamp2CST
// Convert (Unix timestamp) to (CST: UTC+8)
func Timestamp2CST(timestamp string) string {
	return time.Unix(AtoI64(timestamp), 0).
		In(cstLocation()).
		Format(time.DateTime)
}

func cstLocation() *time.Location {
	cstLoc, err := time.LoadLocation(TimeLocationAsiaShanghai)
	if err != nil {
		stdlog.ErrorF("LoadLocation: %s, error: %s", TimeLocationAsiaShanghai, err.Error())
		return nil
	}
	return cstLoc
}

func TimeSecond(sec int) time.Duration {
	return time.Duration(sec) * time.Second
}

func MeasureTime(dur time.Duration) string {
	var (
		appr = float64(dur)
		meas = "ns"
	)
	switch {
	case dur <= 0:
		return ""
	case dur <= time.Microsecond:
		appr /= float64(time.Nanosecond)

	case dur <= time.Millisecond:
		appr /= float64(time.Microsecond)
		meas = "µs"

	case dur <= time.Second:
		appr /= float64(time.Millisecond)
		meas = "ms"

	case dur <= time.Minute:
		appr /= float64(time.Second)
		meas = "s"

	case dur <= time.Hour:
		appr /= float64(time.Minute)
		meas = "min"

	case dur <= DurationDay:
		appr /= float64(time.Hour)
		meas = "hour"

	case dur <= DurationYear:
		appr /= float64(DurationDay)
		meas = "day"

	case dur <= DurationCentury:
		appr /= float64(DurationYear)
		meas = "year"

	default:
		stdlog.WarnF("too long time: %d, do not convert", dur)
		return I64toa(int64(dur)) + meas
	}

	return F64toa(appr) + meas
}

func ToTimestamp(t time.Time) string {
	if t != _zeroTime {
		return I64toa(timestamp.FromTime(t))
	}
	return ""
}

func FromTimestamp(s string) time.Time {
	return timestamp.Time(AtoI64(s))
}

var _zeroTime = time.Time{}

func HumanTimeDesc(d time.Duration) string {
	if ns := d.Nanoseconds(); ns < 1 {
		return "Less than a nanosecond"
	} else if ns < 1000 {
		return fmt.Sprintf("%d nanoseconds", ns)
	} else if mcs := d.Microseconds(); mcs < 1000 {
		return fmt.Sprintf("%d microseconds", mcs)
	} else if mls := d.Milliseconds(); mls < 1000 {
		return fmt.Sprintf("%d ms", mls)
	}
	return unit.HumanDuration(d)
}
