package ms

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const ms float64 = 1
const s = ms * 1000
const m = s * 60
const h = m * 60
const d = h * 24
const w = d * 7
const y = d * 365.25

type parseError struct {
	s string
}

func (e *parseError) Error() string {
	return e.s
}

// Parse string to float64
func Parse(str string) (float64, error) {
	chars := []rune(str)
	if len(chars) > 100 {
		return 0, &parseError{"too much bytes"}
	}

	r, _ := regexp.Compile(`(?i)^((?:\d+)?\-?\d?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|weeks?|w|years?|yrs?|y)?$`)
	match := r.FindStringSubmatch(str)
	if len(match) < 2 {
		return 0, &parseError{"string not in pattern"}
	}

	n, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, &parseError{"string does not contain number"}
	}

	t := "ms"
	if match[2] != "" {
		t = strings.ToLower(match[2])
	}

	switch t {
	case "years", "year", "yrs", "yr", "y":
		return n * y, nil
	case "weeks", "week", "w":
		return n * w, nil
	case "days", "day", "d":
		return n * d, nil
	case "hours", "hour", "hrs", "hr", "h":
		return n * h, nil
	case "minutes", "minute", "mins", "min", "m":
		return n * m, nil
	case "seconds", "second", "secs", "sec", "s":
		return n * s, nil
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return n, nil
	}

	return 0, &parseError{"string can not be parse, bad pattern"}
}

// FmtShort short format
func FmtShort(ms float64) string {
	msAbs := math.Abs(ms)
	if msAbs >= d {
		return strconv.FormatInt(roundFloat64ToInt64(ms/d), 10) + "d"
	}
	if msAbs >= h {
		return strconv.FormatInt(roundFloat64ToInt64(ms/h), 10) + "h"
	}
	if msAbs >= m {
		return strconv.FormatInt(roundFloat64ToInt64(ms/m), 10) + "m"
	}
	if msAbs >= s {
		return strconv.FormatInt(roundFloat64ToInt64(ms/s), 10) + "s"
	}
	return strconv.FormatInt(roundFloat64ToInt64(ms), 10) + "ms"
}

func roundFloat64ToInt64(num float64) int64 {
	return int64(math.Round(num))
}

// FmtLong long format
func FmtLong(val float64) string {
	msAbs := math.Abs(val)
	if msAbs >= d {
		return plural(val, msAbs, d, "day")
	}
	if msAbs >= h {
		return plural(val, msAbs, h, "hour")
	}
	if msAbs >= m {
		return plural(val, msAbs, m, "minute")
	}
	if msAbs >= s {
		return plural(val, msAbs, s, "second")
	}
	return strconv.FormatInt(roundFloat64ToInt64(val), 10) + " ms"
}

func plural(val float64, msAbs float64, n float64, name string) string {
	isPlural := msAbs >= n*1.5
	if isPlural {
		return strconv.FormatInt(roundFloat64ToInt64(val/n), 10) + " " + name + "s"
	}
	return strconv.FormatInt(roundFloat64ToInt64(val/n), 10) + " " + name
}
