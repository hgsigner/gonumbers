package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToCurrency(t *testing.T) {

	a := assert.New(t)

	a.Equal("$1,234,567,890.50",
		NumberToCurrency(1234567890.50))

	a.Equal("$1,234,567,890.51",
		NumberToCurrency(1234567890.506, "precision:2"))

	a.Equal("$1,234,567,890.00",
		NumberToCurrency(1234567890, "precision:2", "unit:$", "separator:."))

	a.Equal("CAD$1,234,567,890.506",
		NumberToCurrency(1234567890.506, "precision:3", "unit:CAD$", "separator:.", "delimiter:,"))

	a.Equal("$1234567890.50",
		NumberToCurrency(1234567890.50, "precision:2", "unit:$", "separator:.", "delimiter:"))

	a.Equal("$1.234.567.890,51",
		NumberToCurrency(1234567890.506, "precision:2", "unit:$", "separator:,", "delimiter:."))

}
