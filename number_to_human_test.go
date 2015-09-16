package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
	"github.com/stretchr/testify/assert"
)

func Test_ToHuman(t *testing.T) {

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
		nth := gonumbers.ToHuman()
		if t.addPrecision {
			nth.Options(gonumbers.Precision(t.precision))
		}
		if t.addSeparator {
			nth.Options(gonumbers.Separator(t.separator))
		}
		a.Equal(t.out, nth.Perform(t.in))
	}

}

func ExampleToHuman() {
	nth1 := gonumbers.ToHuman()
	fmt.Println(nth1.Perform(123))

	nth2 := gonumbers.ToHuman()
	fmt.Println(nth2.Perform(1234))

	nth3 := gonumbers.ToHuman()
	fmt.Println(nth3.Perform(12345))

	nth4 := gonumbers.ToHuman()
	fmt.Println(nth4.Perform(1234567))

	nth5 := gonumbers.ToHuman()
	fmt.Println(nth5.Perform(1234567890))

	nth6 := gonumbers.ToHuman()
	fmt.Println(nth6.Perform(1234567890123))

	nth7 := gonumbers.ToHuman()
	fmt.Println(nth7.Perform(1234567890123456))

	nth8 := gonumbers.ToHuman()
	nth8.Options(gonumbers.Precision(2))
	fmt.Println(nth8.Perform(489939))

	nth9 := gonumbers.ToHuman()
	nth9.Options(gonumbers.Precision(4))
	fmt.Println(nth9.Perform(489939))

	nth10 := gonumbers.ToHuman()
	nth10.Options(gonumbers.Precision(4), gonumbers.Separator(","))
	fmt.Println(nth10.Perform(489939))

	// Output:
	// 123
	// 1.23 Thousand
	// 12.3 Thousand
	// 1.23 Million
	// 1.23 Billion
	// 1.23 Trillion
	// 1.23 Quadrillion
	// 490 Thousand
	// 489.9 Thousand
	// 489,9 Thousand
}
