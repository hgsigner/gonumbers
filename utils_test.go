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
			in:        "1234567890",
			precision: 3,
			out:       []string{"1", "234", "567", "890"},
		},
		{
			in:        "123456",
			precision: 3,
			out:       []string{"123", "456"},
		},
		{
			in:        "1234",
			precision: 3,
			out:       []string{"1", "234"},
		},
		{
			in:        "123",
			precision: 3,
			out:       []string{"123"},
		},
		{
			in:        "1234567890",
			precision: 4,
			out:       []string{"12", "3456", "7890"},
		},
		{
			in:        "123456",
			precision: 4,
			out:       []string{"12", "3456"},
		},
		{
			in:        "1234",
			precision: 4,
			out:       []string{"1234"},
		},
		{
			in:        "123",
			precision: 4,
			out:       []string{"0.0123"},
		},
		{
			in:        "123",
			precision: 5,
			out:       []string{"0.00123"},
		},
	}

	for _, item := range tests {
		a.Equal(item.out, group_in_thousands(item.in, item.precision))
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
