package neralie

import (
	"fmt"
	"time"
)

type Neralie int64

func (n Neralie) String() string {
	ms := time.Duration(n).Milliseconds()
	v := fmt.Sprintf("%.6f", float64(ms)/8640.0/10000.0)
	return fmt.Sprintf("%s:%s", v[2:5], v[5:8])
}

func ToTime(n Neralie) time.Time {
	return bod(time.Now()).Add(time.Duration(n))
}

func FromTime(t time.Time) Neralie {
	return Neralie(t.Sub(bod(t)))
}

func bod(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
