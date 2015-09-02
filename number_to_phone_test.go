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
