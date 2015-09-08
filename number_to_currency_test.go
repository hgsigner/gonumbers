package gonumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsNTC = []struct {
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

func Example_NumberToCurrency() {
	ntc1 := NewNumberToCurrency()
	res1 := ntc1.Perform(1234567890.50)
	fmt.Println(res1)

	ntc2 := NewNumberToCurrency()
	ntc2.Options(Precision(2))
	res2 := ntc2.Perform(1234567890.506)
	fmt.Println(res2)

	ntc3 := NewNumberToCurrency()
	ntc3.Options(Precision(2), Unit("$"), Separator("."))
	res3 := ntc3.Perform(1234567890)
	fmt.Println(res3)

	ntc4 := NewNumberToCurrency()
	ntc4.Options(Precision(3), Unit("CAD$"), Separator("."), Delimiter(","))
	res4 := ntc4.Perform(1234567890.506)
	fmt.Println(res4)

	ntc5 := NewNumberToCurrency()
	ntc5.Options(Precision(2), Separator("."), Delimiter(""))
	res5 := ntc5.Perform(1234567890.50)
	fmt.Println(res5)

	ntc6 := NewNumberToCurrency()
	ntc6.Options(Precision(2), Separator(","), Delimiter("."))
	res6 := ntc6.Perform(1234567890.506)
	fmt.Println(res6)

	// Output:
	// $1,234,567,890.50
	// $1,234,567,890.51
	// $1,234,567,890.00
	// CAD$1,234,567,890.506
	// $1234567890.50
	// $1.234.567.890,51
}

func Test_NumberToCurrency(t *testing.T) {

	a := assert.New(t)

	for _, t := range testsNTC {
		ntc := NewNumberToCurrency()

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
