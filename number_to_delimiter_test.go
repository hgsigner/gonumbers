package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
)

func Test_ToDelimiter(t *testing.T) {

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

	for _, test := range tests {
		ntd := gonumbers.ToDelimiter()
		if test.addSeparator {
			ntd.Options(gonumbers.Separator(test.separator))
		}
		if test.addDelimiter {
			ntd.Options(gonumbers.Delimiter(test.delimiter))
		}
		assert(t, test.out, ntd.Perform(test.in))
	}

}

func ExampleToDelimiter() {
	ntd1 := gonumbers.ToDelimiter()
	fmt.Println(ntd1.Perform(1234567890))

	ntd2 := gonumbers.ToDelimiter()
	ntd2.Options(gonumbers.Separator("."), gonumbers.Delimiter(","))
	fmt.Println(ntd2.Perform(1234567890.506))

	ntd3 := gonumbers.ToDelimiter()
	ntd3.Options(gonumbers.Separator("-"), gonumbers.Delimiter(":"))
	fmt.Println(ntd3.Perform(1234567890.34))

	ntd4 := gonumbers.ToDelimiter()
	ntd4.Options(gonumbers.Separator("."), gonumbers.Delimiter(""))
	fmt.Println(ntd4.Perform(1234567890.34))

	ntd5 := gonumbers.ToDelimiter()
	ntd5.Options(gonumbers.Separator(""), gonumbers.Delimiter(""))
	fmt.Println(ntd5.Perform(1234567890.34))

	// Output:
	// 1,234,567,890
	// 1,234,567,890.506
	// 1:234:567:890-34
	// 1234567890.34
	// 123456789034
}
