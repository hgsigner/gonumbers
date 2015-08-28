package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Currency(t *testing.T) {

	a := assert.New(t)

	c := NumberToCurrency(1234567890.50, 2, "$")

	a.Equal("$1,234,567,890.50", c)
}
