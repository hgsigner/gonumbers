package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToHuman(t *testing.T) {

	a := assert.New(t)

	tests := []struct {
		in                         float64
		out                        string
		precision                  int
		separator                  string
		addPrecision, addSeparator bool
	}{
		{in: 123, out: "123"},
		{in: 1234, out: "1.23 Thousand"},
		{in: 12345, out: "12.3 Thousand"},
		{in: 1234567, out: "1.23 Million"},
		{in: 1234567890, out: "1.23 Billion"},
		{in: 1234567890123, out: "1.23 Trillion"},
		{in: 1234567890123456, out: "1.23 Quadrillion"},
		{
			in:           489939,
			addPrecision: true,
			precision:    2,
			out:          "490 Thousand",
		},
		{
			in:           489939,
			addPrecision: true,
			precision:    4,
			out:          "489.9 Thousand",
		},
		{
			in:           489939,
			addPrecision: true,
			precision:    4,
			addSeparator: true,
			separator:    ",",
			out:          "489,9 Thousand",
		},
	}

	for _, t := range tests {
		nth := new(NumberToHuman)
		if t.addPrecision {
			nth.Options(nth.Precision(t.precision))
		}
		if t.addSeparator {
			nth.Options(nth.Separator(t.separator))
		}
		a.Equal(t.out, nth.Perform(t.in))
	}

}
