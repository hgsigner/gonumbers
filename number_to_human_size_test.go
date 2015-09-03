package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToHumanSize(t *testing.T) {
	a := assert.New(t)

	t1, err1 := NumberToHumanSize(123)
	a.NoError(err1)
	a.Equal("123 Bytes", t1)

}
