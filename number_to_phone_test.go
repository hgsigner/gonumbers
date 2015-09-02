package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToPhone(t *testing.T) {
	a := assert.New(t)

	t_wa1, e_wa1 := NumberToPhone("good")
	a.Error(e_wa1)
	a.Contains(e_wa1.Error(), "The value should an integer.")
	a.Equal("", t_wa1)

	t_wa2, e_wa2 := NumberToPhone("123abc456")
	a.Error(e_wa2)
	a.Contains(e_wa2.Error(), "The value should an integer.")
	a.Equal("", t_wa2)

	t1, e1 := NumberToPhone(5551234)
	a.NoError(e1)
	a.Equal("555-1234", t1)

	t2, e2 := NumberToPhone(1235551234)
	a.NoError(e2)
	a.Equal("123-555-1234", t2)

	t3, e3 := NumberToPhone(1235551234, "area_code:true")
	a.NoError(e3)
	a.Equal("(123) 555-1234", t3)

	t4, e4 := NumberToPhone(1235551234, "area_code:true", "country_code:1")
	a.NoError(e4)
	a.Equal("+1(123) 555-1234", t4)

	t5, e5 := NumberToPhone(1235551234, "country_code:1")
	a.NoError(e5)
	a.Equal("+1-123-555-1234", t5)

	t6, e6 := NumberToPhone(5551234, "area_code:true")
	a.NoError(e6)
	a.Equal("555-1234", t6)

	t7, e7 := NumberToPhone(1234555556789, "area_code:true", "country_code:1", "extension:4545")
	a.NoError(e7)
	a.Equal("+1(123455) 555-6789 x 4545", t7)

	t8, e8 := NumberToPhone(1234555556789, "area_code:true", "country_code:1", "extension:4545", "digits_size:4")
	a.NoError(e8)
	a.Equal("+1(12345) 5555-6789 x 4545", t8)

	t9, e9 := NumberToPhone(1234555556789, "area_code:true", "country_code:1", "extension:4545", "digits_size:5")
	a.NoError(e9)
	a.Equal("+1(1234) 55555-6789 x 4545", t9)

	t10, e10 := NumberToPhone(1234555556789, "area_code:true", "country_code:55", "digits_size:5", "delimiter:,")
	a.NoError(e10)
	a.Equal("+55(1234) 55555,6789", t10)
}

func Test_SplitToPhoneFormat(t *testing.T) {
	a := assert.New(t)

	a.Equal([]string{"555", "1234"}, split_to_phone_format("5551234", 3))
	a.Equal([]string{"5", "555", "1234"}, split_to_phone_format("55551234", 3))
	a.Equal([]string{"5555", "1234"}, split_to_phone_format("55551234", 4))
	a.Equal([]string{"123", "555", "1234"}, split_to_phone_format("1235551234", 3))
	a.Equal([]string{"123123", "555", "1234"}, split_to_phone_format("1231235551234", 3))
	a.Equal([]string{"12312", "3555", "1234"}, split_to_phone_format("1231235551234", 4))
}
