package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToDelimiter(t *testing.T) {

	a := assert.New(t)

	a.Equal("1,234,567,890",
		NumberToDelimiter(1234567890))

	a.Equal("1,234,567,890.506",
		NumberToDelimiter(1234567890.506, "separator:.", "delimiter:,"))

	a.Equal("1:234:567:890-34",
		NumberToDelimiter(1234567890.34, "separator:-", "delimiter::"))

	a.Equal("1234567890.34",
		NumberToDelimiter(1234567890.34, "separator:.", "delimiter:"))
}
