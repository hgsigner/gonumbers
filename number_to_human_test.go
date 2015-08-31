package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToHuman(t *testing.T) {

	a := assert.New(t)

	//Test with defaults

	a.Equal("1.234 Thousand", NumberToHuman(1234))
	a.Equal("12.3 Thousand", NumberToHuman(12345))
	a.Equal("123.4 Thousand", NumberToHuman(123456))

	a.Equal("1.234 Million", NumberToHuman(1234567))
	a.Equal("12.3 Million", NumberToHuman(12345678))
	a.Equal("123.4 Million", NumberToHuman(123456789))

	a.Equal("123 Million", NumberToHuman(123000000))

	//Custom

	a.Equal("1,234 Thousand", NumberToHuman(1234, "separator:,"))
	a.Equal("1234", NumberToHuman(1234, "precision:4"))
	a.Equal("0.01234", NumberToHuman(1234, "precision:5"))
	a.Equal("0,01234", NumberToHuman(1234, "precision:5", "separator:,"))
	// a.Equal("12.3 Thousand", NumberToHuman(12345))
	// a.Equal("123.4 Thousand", NumberToHuman(123456))

	// a.Equal("1.234 Million", NumberToHuman(1234567))
	// a.Equal("12.3 Million", NumberToHuman(12345678))
	// a.Equal("123.4 Million", NumberToHuman(123456789))

	// a.Equal("123 Million", NumberToHuman(123000000))

}
