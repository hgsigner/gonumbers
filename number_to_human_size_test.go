package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToHumanSize(t *testing.T) {
	a := assert.New(t)

	nths1 := NewNumberToHumanSize()
	nths1.Options(Prefix("ahhaha"))
	nths1_final, err1 := nths1.Perform(123)
	a.Error(err1)
	a.Contains(err1.Error(), "Prefix must be binary or si.")
	a.Equal("", nths1_final)

	tests := []struct {
		in           float64
		addPrecision bool
		precision    int
		addSeparator bool
		separator    string
		addDelimiter bool
		delimiter    string
		addPrefix    bool
		prefix       string
		out          string
	}{
		{
			in:  333,
			out: "333 Bytes",
		},
		{
			in:  1234,
			out: "1.21 KB",
		},
		{
			in:  1234567890,
			out: "1.15 GB",
		},
		{
			in:           1234567,
			addPrecision: true,
			precision:    2,
			out:          "1.2 MB",
		},
		{
			in:           524288000,
			addPrecision: true,
			precision:    5,
			out:          "500 MB",
		},
		{
			in:           1234567,
			addPrecision: true,
			precision:    2,
			addSeparator: true,
			separator:    ",",
			out:          "1,2 MB",
		},
	}

	for _, t := range tests {
		nths := NewNumberToHumanSize()
		if t.addPrecision {
			nths.Options(Precision(t.precision))
		}
		if t.addSeparator {
			nths.Options(Separator(t.separator))
		}
		if t.addDelimiter {
			nths.Options(Delimiter(t.delimiter))
		}
		if t.addPrefix {
			nths.Options(Prefix(t.prefix))
		}
		nths_final, err := nths.Perform(t.in)
		a.NoError(err)
		a.Equal(t.out, nths_final)
	}

}
