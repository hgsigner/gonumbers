package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToDelimiter(t *testing.T) {

	a := assert.New(t)

	ntd1 := new(NumberToDelimiter)
	ntd1_final := ntd1.Perform(1234567890)
	a.Equal("1,234,567,890", ntd1_final)

	ntd2 := new(NumberToDelimiter)
	ntd2.Options(ntd2.Separator("."), ntd2.Delimiter(","))
	ntd2_final := ntd2.Perform(1234567890.506)
	a.Equal("1,234,567,890.506", ntd2_final)

	ntd3 := new(NumberToDelimiter)
	ntd3.Options(ntd3.Separator("-"), ntd3.Delimiter(":"))
	ntd3_final := ntd3.Perform(1234567890.34)
	a.Equal("1:234:567:890-34", ntd3_final)

	ntd4 := new(NumberToDelimiter)
	ntd4.Options(ntd4.Separator("."), ntd4.Delimiter(""))
	ntd4_final := ntd4.Perform(1234567890.34)
	a.Equal("1234567890.34", ntd4_final)

	ntd5 := new(NumberToDelimiter)
	ntd5.Options(ntd5.Separator(""), ntd5.Delimiter(""))
	ntd5_final := ntd5.Perform(1234567890.34)
	a.Equal("123456789034", ntd5_final)

}
