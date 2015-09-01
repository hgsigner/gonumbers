package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToHuman(t *testing.T) {

	a := assert.New(t)

	//Test with defaults

	a.Equal("123", NumberToHuman(123))
	a.Equal("1.23 Thousand", NumberToHuman(1234))
	a.Equal("12.3 Thousand", NumberToHuman(12345))
	a.Equal("1.23 Million", NumberToHuman(1234567))
	a.Equal("1.23 Billion", NumberToHuman(1234567890))
	a.Equal("1.23 Trillion", NumberToHuman(1234567890123))
	a.Equal("1.23 Quadrillion", NumberToHuman(1234567890123456))

	//Custom

	a.Equal("490 Thousand", NumberToHuman(489939, "precision:2"))
	a.Equal("489.9 Thousand", NumberToHuman(489939, "precision:4"))
	a.Equal("489,9 Thousand", NumberToHuman(489939, "precision:4", "separator:,"))

}
