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
			out:       []string{"123"},
		},
	}

	for _, item := range tests {
		a.Equal(item.out, group_in_thousands(item.in, item.precision))
	}

}
