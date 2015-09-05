package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GroupInThousands(t *testing.T) {

	a := assert.New(t)

	tests := []struct {
		in        string
		precision int
		out       []string
	}{
		{
			in:  "1234567890",
			out: []string{"1", "234", "567", "890"},
		},
		{
			in:  "123456",
			out: []string{"123", "456"},
		},
		{
			in:  "1234",
			out: []string{"1", "234"},
		},
		{
			in:  "123",
			out: []string{"123"},
		},
	}

	for _, item := range tests {
		a.Equal(item.out, group_in_thousands(item.in))
	}

}

func Test_RoundedNumber(t *testing.T) {
	a := assert.New(t)
	a.Equal("490000", rounded_number(489939, 2))
}

func Test_RoundedWithPrecision(t *testing.T) {
	a := assert.New(t)
	a.Equal(1.2, round_with_precision(1.2000000000000002, 3))
	a.Equal(1.27, round_with_precision(1.267800001, 2))
	a.Equal(1.268, round_with_precision(1.267800001, 3))
}

func Test_DigtsCount(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, digits_count(123456))
	a.Equal(10, digits_count(1234567890))
}
