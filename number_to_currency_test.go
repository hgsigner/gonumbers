package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToCurrency(t *testing.T) {

	a := assert.New(t)

	tests := []struct {
		in           float64
		addPrecision bool
		precision    int
		addUnit      bool
		unit         string
		addSeparator bool
		separator    string
		addDelimiter bool
		delimiter    string
		out          string
	}{
		{in: 1234567890.50, out: "$1,234,567,890.50"},
		{
			in:           1234567890.506,
			addPrecision: true,
			precision:    2,
			out:          "$1,234,567,890.51",
		},
		{
			in:           1234567890,
			addPrecision: true,
			precision:    2,
			addUnit:      true,
			unit:         "$",
			addSeparator: true,
			separator:    ".",
			out:          "$1,234,567,890.00",
		},
		{
			in:           1234567890.506,
			addPrecision: true,
			precision:    3,
			addUnit:      true,
			unit:         "CAD$",
			addSeparator: true,
			separator:    ".",
			addDelimiter: true,
			delimiter:    ",",
			out:          "CAD$1,234,567,890.506",
		},
		{
			in:           1234567890.50,
			addPrecision: true,
			precision:    2,
			addSeparator: true,
			separator:    ".",
			addDelimiter: true,
			delimiter:    "",
			out:          "$1234567890.50",
		},
		{
			in:           1234567890.506,
			addPrecision: true,
			precision:    2,
			addSeparator: true,
			separator:    ",",
			addDelimiter: true,
			delimiter:    ".",
			out:          "$1.234.567.890,51",
		},
	}

	for _, t := range tests {
		ntc := &NumberToCurrency{}

		if t.addPrecision {
			ntc.Options(Precision(t.precision))
		}

		if t.addUnit {
			ntc.Options(Unit(t.unit))
		}

		if t.addSeparator {
			ntc.Options(Separator(t.separator))
		}

		if t.addDelimiter {
			ntc.Options(Delimiter(t.delimiter))
		}

		a.Equal(t.out, ntc.Perform(t.in))
	}

}
