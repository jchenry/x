package neralie_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jchenry/x/time/neralie"
)

func TestFromTime(t *testing.T) {

	tests := [][]string{
		{"250:000", "6:00"},
		{"500:000", "12:00"},
		{"750:000", "18:00"},
	}
	for i := range tests {
		y, m, d := time.Now().Date()
		dt, _ := time.Parse("15:04", tests[i][1])
		h, m2, s := dt.Clock()
		n := neralie.FromTime(time.Date(y, m, d, h, m2, s, 0, time.UTC))
		expected := tests[i][0]
		if !strings.EqualFold(n.String(), expected) {
			fmt.Printf("%s != %s\n", expected, n.String())
			t.Fail()
		}
	}
}

func TestToTime(t *testing.T) {
	tests := [][]string{
		{"250:000", "6:00"},
		{"500:000", "12:00"},
		{"750:000", "18:00"},
	}

	for i := range tests {
		y, m, d := time.Now().Date()
		dt, _ := time.Parse("15:04", tests[i][1])
		h, m2, s := dt.Clock()
		d1 := time.Date(y, m, d, h, m2, s, 0, time.UTC)
		n := neralie.FromTime(d1)
		d2 := neralie.ToTime(n)

		if !d1.Equal(d2) {
			fmt.Printf("%s != %s\n", d1, d2)
			t.Fail()
		}

	}
}
