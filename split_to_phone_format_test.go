package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SplitToPhoneFormat(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		val       string
		precision int
		out       []string
	}{
		{
			val:       "5551234",
			precision: 3,
			out:       []string{"555", "1234"},
		},
		{
			val:       "55551234",
			precision: 3,
			out:       []string{"5", "555", "1234"},
		},
		{
			val:       "55551234",
			precision: 4,
			out:       []string{"5555", "1234"},
		},
		{
			val:       "1235551234",
			precision: 3,
			out:       []string{"123", "555", "1234"},
		},
		{
			val:       "1231235551234",
			precision: 3,
			out:       []string{"123123", "555", "1234"},
		},
		{
			val:       "1231235551234",
			precision: 4,
			out:       []string{"12312", "3555", "1234"},
		},
	}

	for _, t := range tests {
		a.Equal(t.out, split_to_phone_format(t.val, t.precision))
	}

}
