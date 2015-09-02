package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToPercentage(t *testing.T) {

	a := assert.New(t)
	a.Equal("100.000%", NumberToPercentage(100))
	a.Equal("98.000%", NumberToPercentage("98"))
	a.Equal("100%", NumberToPercentage(100, "precision:0"))
	a.Equal("1,000.000%", NumberToPercentage(1000))
	a.Equal("1.000,000%", NumberToPercentage(1000, "delimiter:.", "separator:,"))
	a.Equal("302.24399%", NumberToPercentage(302.24398923423, "precision:5"))
	a.Equal("0.000%", NumberToPercentage("hahahah"))

}
