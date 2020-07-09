package arvelie_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jchenry/x/time/arvelie"
)

func TestFromDate(t *testing.T) {
	tests := [][]string{
		{"02A01", "2002-01-01"},
		{"13B12", "2013-01-26"},
		{"24C01", "2024-01-29"},
		{"01D07", "2001-02-18"},
		{"02E07", "2002-03-04"},
		{"03+01", "2003-12-31"},
	}

	for i := range tests {
		dt, _ := time.Parse("2006-01-02", tests[i][1])
		a := arvelie.FromDate(dt)
		expected := tests[i][0]
		if !strings.EqualFold(string(a), expected) {
			fmt.Printf("%v != %v\n", expected, a)
			t.Fail()
		}

	}
}

func TestToDate(t *testing.T) {
	tests := [][]string{
		{"02A01", "2002-01-01"},
		{"13B12", "2013-01-26"},
		{"24C01", "2024-01-29"},
		{"01D07", "2001-02-18"},
		{"02E07", "2002-03-04"},
		{"03+01", "2003-12-31"},
	}

	for i := range tests {
		d1, _ := time.Parse("2006-01-02", tests[i][1])
		dt := arvelie.ToDate(arvelie.Arvelie(tests[i][0]))
		if !d1.Equal(dt) {
			fmt.Printf("%s != %s\n", d1, dt)
			t.Fail()
		}
	}
}
