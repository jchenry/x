package arvelie

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Arvelie string

func (a *Arvelie) isValid() bool {
	if a != nil {
		return strings.EqualFold(string(*a), string(FromDate(ToDate(*a))))
	}
	return false
}

func ToDate(a Arvelie) time.Time {
	y := string(a)[0:2]
	m := string(a)[2:3]
	d, _ := strconv.Atoi(string(a)[3:5])

	var mon int
	if m == "+" {
		mon = 26
	} else {
		mon = (int(m[0]) - 65)
	}

	doty := (math.Floor(float64(mon)*14) + math.Floor(float64(d)) - 1)
	yr, _ := strconv.Atoi(fmt.Sprintf("20%s", y))
	return time.Date(yr, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(doty))
}

func FromDate(date time.Time) Arvelie {
	y := date.Format("06")
	doty := date.YearDay()

	var m string
	if doty == 365 || doty == 366 {
		m = "+"
	} else {
		m = strings.ToUpper(string([]byte{byte(97 + math.Floor(float64(doty/14)))}))
	}

	var d string
	switch doty {
	case 365:
		d = fmt.Sprintf("%02d", 1)
		break
	case 366:
		d = fmt.Sprintf("%02d", 2)
		break
	default:
		d = fmt.Sprintf("%02d", (doty % 14))
	}

	return Arvelie(fmt.Sprintf("%s%s%s", y, m, d))
}
