package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToDelimiter(t *testing.T) {

	a := assert.New(t)

	tests := []struct {
		in           float64
		addSeparator bool
		separator    string
		addDelimiter bool
		delimiter    string
		out          string
	}{
		{in: 1234567890, out: "1,234,567,890"},
		{
			in:           1234567890.506,
			addSeparator: true,
			separator:    ".",
			addDelimiter: true,
			delimiter:    ",",
			out:          "1,234,567,890.506",
		},
		{
			in:           1234567890.34,
			addSeparator: true,
			separator:    "-",
			addDelimiter: true,
			delimiter:    ":",
			out:          "1:234:567:890-34",
		},
		{
			in:           1234567890.34,
			addSeparator: true,
			separator:    ".",
			addDelimiter: true,
			delimiter:    "",
			out:          "1234567890.34",
		},
		{
			in:           1234567890.34,
			addSeparator: true,
			separator:    "",
			addDelimiter: true,
			delimiter:    "",
			out:          "123456789034",
		},
	}

	for _, t := range tests {
		ntd := new(NumberToDelimiter)
		if t.addSeparator {
			ntd.Options(Separator(t.separator))
		}
		if t.addDelimiter {
			ntd.Options(Delimiter(t.delimiter))
		}
		a.Equal(t.out, ntd.Perform(t.in))
	}

}
