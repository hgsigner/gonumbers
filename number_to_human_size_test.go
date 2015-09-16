package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
	"github.com/stretchr/testify/assert"
)

func Test_ToHumanSize(t *testing.T) {
	a := assert.New(t)

	nths1 := gonumbers.ToHumanSize()
	nths1.Options(gonumbers.Prefix("ahhaha"))
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
		nths := gonumbers.ToHumanSize()
		if t.addPrecision {
			nths.Options(gonumbers.Precision(t.precision))
		}
		if t.addSeparator {
			nths.Options(gonumbers.Separator(t.separator))
		}
		if t.addDelimiter {
			nths.Options(gonumbers.Delimiter(t.delimiter))
		}
		if t.addPrefix {
			nths.Options(gonumbers.Prefix(t.prefix))
		}
		nths_final, err := nths.Perform(t.in)
		a.NoError(err)
		a.Equal(t.out, nths_final)
	}

}

func ExampleToHumanSize() {
	nths1 := gonumbers.ToHumanSize()
	res1, err1 := nths1.Perform(333)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res1)

	nths2 := gonumbers.ToHumanSize()
	res2, err2 := nths2.Perform(1234)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(res2)

	nths3 := gonumbers.ToHumanSize()
	res3, err3 := nths3.Perform(1234567890)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(res3)

	nths4 := gonumbers.ToHumanSize()
	nths4.Options(gonumbers.Precision(2))
	res4, err4 := nths4.Perform(1234567)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(res4)

	nths5 := gonumbers.ToHumanSize()
	nths5.Options(gonumbers.Precision(2), gonumbers.Separator(","))
	res5, err5 := nths5.Perform(1234567)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(res5)

	nths6 := gonumbers.ToHumanSize()
	nths6.Options(gonumbers.Precision(5))
	res6, err6 := nths6.Perform(524288000)
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println(res6)

	// Output:
	// 333 Bytes
	// 1.21 KB
	// 1.15 GB
	// 1.2 MB
	// 1,2 MB
	// 500 MB
}
