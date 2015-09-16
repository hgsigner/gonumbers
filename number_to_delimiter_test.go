package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
	"github.com/stretchr/testify/assert"
)

func Test_ToDelimiter(t *testing.T) {

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
		ntd := gonumbers.ToDelimiter()
		if t.addSeparator {
			ntd.Options(gonumbers.Separator(t.separator))
		}
		if t.addDelimiter {
			ntd.Options(gonumbers.Delimiter(t.delimiter))
		}
		a.Equal(t.out, ntd.Perform(t.in))
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
