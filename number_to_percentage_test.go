package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
)

func Test_ToPercentage(t *testing.T) {

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

	for _, test := range tests {
		ntp := gonumbers.ToPercentage()

		if test.addPrecision {
			ntp.Options(gonumbers.Precision(test.precision))
		}

		if test.addSeparator {
			ntp.Options(gonumbers.Separator(test.separator))
		}

		if test.addDelimiter {
			ntp.Options(gonumbers.Delimiter(test.delimiter))
		}

		assert(t, test.out, ntp.Perform(test.in))
	}

}

func ExampleToPercentage() {
	ntp1 := gonumbers.ToPercentage()
	fmt.Println(ntp1.Perform(100))

	ntp2 := gonumbers.ToPercentage()
	fmt.Println(ntp2.Perform("98"))

	ntp3 := gonumbers.ToPercentage()
	ntp3.Options(gonumbers.Precision(0))
	fmt.Println(ntp3.Perform(100))

	ntp4 := gonumbers.ToPercentage()
	fmt.Println(ntp4.Perform(1000))

	ntp5 := gonumbers.ToPercentage()
	ntp5.Options(gonumbers.Delimiter("."), gonumbers.Separator(","))
	fmt.Println(ntp5.Perform(1000))

	ntp6 := gonumbers.ToPercentage()
	ntp6.Options(gonumbers.Precision(5))
	fmt.Println(ntp6.Perform(302.24398923423))

	ntp7 := gonumbers.ToPercentage()
	fmt.Println(ntp7.Perform("hahahah"))

	// Output:
	// 100.000%
	// 98.000%
	// 100%
	// 1,000.000%
	// 1.000,000%
	// 302.24399%
	// 0.000%
}
