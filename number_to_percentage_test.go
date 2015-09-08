package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToPercentage(t *testing.T) {

	a := assert.New(t)

	tests := []struct {
		in                                       interface{}
		out                                      string
		precision                                int
		separator, delimiter                     string
		addPrecision, addSeparator, addDelimiter bool
	}{
		{in: 100, out: "100.000%"},
		{in: "98", out: "98.000%"},
		{
			in:           100,
			addPrecision: true,
			precision:    0,
			out:          "100%",
		},
		{in: 1000, out: "1,000.000%"},
		{
			in:           1000,
			addDelimiter: true,
			delimiter:    ".",
			addSeparator: true,
			separator:    ",",
			out:          "1.000,000%",
		},
		{
			in:           302.24398923423,
			addPrecision: true,
			precision:    5,
			out:          "302.24399%",
		},
		{in: "hahahah", out: "0.000%"},
	}

	for _, t := range tests {
		ntp := NewNumberToPercentage()

		if t.addPrecision {
			ntp.Options(Precision(t.precision))
		}

		if t.addSeparator {
			ntp.Options(Separator(t.separator))
		}

		if t.addDelimiter {
			ntp.Options(Delimiter(t.delimiter))
		}

		a.Equal(t.out, ntp.Perform(t.in))
	}

}
