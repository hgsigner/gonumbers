package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToCurrency(t *testing.T) {

	a := assert.New(t)

	ntc2 := new(NumberToCurrency)
	ntc2_final := ntc2.Perform(1234567890.50)
	a.Equal("$1,234,567,890.50", ntc2_final)

	ntc3 := new(NumberToCurrency)
	ntc3.Options(ntc3.Precision(2))
	ntc3_final := ntc3.Perform(1234567890.506)
	a.Equal("$1,234,567,890.51", ntc3_final)

	ntc4 := new(NumberToCurrency)
	ntc4.Options(ntc4.Precision(2), ntc4.Unit("$"), ntc4.Separator("."))
	ntc4_final := ntc4.Perform(1234567890)
	a.Equal("$1,234,567,890.00", ntc4_final)

	ntc5 := new(NumberToCurrency)
	ntc5.Options(ntc5.Precision(3), ntc5.Unit("CAD$"), ntc5.Separator("."), ntc5.Delimiter(","))
	ntc5_final := ntc5.Perform(1234567890.506)
	a.Equal("CAD$1,234,567,890.506", ntc5_final)

	ntc6 := new(NumberToCurrency)
	ntc6.Options(ntc6.Precision(2), ntc5.Separator("."), ntc5.Delimiter(""))
	ntc6_final := ntc6.Perform(1234567890.50)
	a.Equal("$1234567890.50", ntc6_final)

	ntc7 := new(NumberToCurrency)
	ntc7.Options(ntc7.Precision(2), ntc7.Separator(","), ntc7.Delimiter("."))
	ntc7_final := ntc7.Perform(1234567890.506)
	a.Equal("$1.234.567.890,51", ntc7_final)

}
