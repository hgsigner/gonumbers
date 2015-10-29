package gonumbers_test

import (
	"testing"

	"github.com/hgsigner/gonumbers"
)

func Test_GroupInThousands(t *testing.T) {

	//a := assert.New(t)

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
		assert(t, item.out, gonumbers.GroupInThousands(item.in))
	}

}

func Test_RoundedNumber(t *testing.T) {
	//a := assert.New(t)
	assert(t, "490000", gonumbers.RoundedNumber(489939, 2))
}

func Test_RoundedWithPrecision(t *testing.T) {
	//a := assert.New(t)
	assert(t, 1.2, gonumbers.RoundWithPrecision(1.2000000000000002, 3))
	assert(t, 1.27, gonumbers.RoundWithPrecision(1.267800001, 2))
	assert(t, 1.268, gonumbers.RoundWithPrecision(1.267800001, 3))
}

func Test_DigtsCount(t *testing.T) {
	//a := assert.New(t)
	assert(t, 6, gonumbers.DigitsCount(123456))
	assert(t, 10, gonumbers.DigitsCount(1234567890))
}
