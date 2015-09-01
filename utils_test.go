package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupInThousands(t *testing.T) {

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

func TestFuncParams(t *testing.T) {
	a := assert.New(t)

	u1, s1, p1, d1 := func_params()
	a.Equal("$notset$", u1)
	a.Equal("$notset$", s1)
	a.Equal(0, p1)
	a.Equal("$notset$", d1)

	u2, s2, p2, d2 := func_params("separator:.", "precision:20")
	a.Equal("$notset$", u2)
	a.Equal(".", s2)
	a.Equal(20, p2)
	a.Equal("$notset$", d2)

	u3, s3, p3, d3 := func_params("delimiter:,", "precision:3", "separator:.", "unit:$")
	a.Equal("$", u3)
	a.Equal(".", s3)
	a.Equal(3, p3)
	a.Equal(",", d3)

}

func TestRoundedNumber(t *testing.T) {
	a := assert.New(t)

	a.Equal("490000", rounded_number(489939, 2))
}
