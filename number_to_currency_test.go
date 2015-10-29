package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
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

func Test_ToCurrency(t *testing.T) {

	for _, test := range testsNTC {
		ntc := gonumbers.ToCurrency()

		if test.addPrecision {
			ntc.Options(gonumbers.Precision(test.precision))
		}

		if test.addUnit {
			ntc.Options(gonumbers.Unit(test.unit))
		}

		if test.addSeparator {
			ntc.Options(gonumbers.Separator(test.separator))
		}

		if test.addDelimiter {
			ntc.Options(gonumbers.Delimiter(test.delimiter))
		}

		assert(t, test.out, ntc.Perform(test.in))
	}

}

func ExampleToCurrency() {
	ntc1 := gonumbers.ToCurrency()
	res1 := ntc1.Perform(1234567890.50)
	fmt.Println(res1)

	ntc2 := gonumbers.ToCurrency()
	ntc2.Options(gonumbers.Precision(2))
	res2 := ntc2.Perform(1234567890.506)
	fmt.Println(res2)

	ntc3 := gonumbers.ToCurrency()
	ntc3.Options(gonumbers.Precision(2), gonumbers.Unit("$"), gonumbers.Separator("."))
	res3 := ntc3.Perform(1234567890)
	fmt.Println(res3)

	ntc4 := gonumbers.ToCurrency()
	ntc4.Options(gonumbers.Precision(3), gonumbers.Unit("CAD$"), gonumbers.Separator("."), gonumbers.Delimiter(","))
	res4 := ntc4.Perform(1234567890.506)
	fmt.Println(res4)

	ntc5 := gonumbers.ToCurrency()
	ntc5.Options(gonumbers.Precision(2), gonumbers.Separator("."), gonumbers.Delimiter(""))
	res5 := ntc5.Perform(1234567890.50)
	fmt.Println(res5)

	ntc6 := gonumbers.ToCurrency()
	ntc6.Options(gonumbers.Precision(2), gonumbers.Separator(","), gonumbers.Delimiter("."))
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
