package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToHumanSize(t *testing.T) {
	a := assert.New(t)

	tw1, twerr1 := NumberToHumanSize(123, "prefix:hahah")
	a.Error(twerr1)
	a.Contains(twerr1.Error(), "Prefix must be binary or si.")
	a.Equal("", tw1)

	t1, err1 := NumberToHumanSize(123)
	a.NoError(err1)
	a.Equal("123 Bytes", t1)

	t2, err2 := NumberToHumanSize(1234)
	a.NoError(err2)
	a.Equal("1.21 KB", t2)

	t3, err3 := NumberToHumanSize(1234567890)
	a.NoError(err3)
	a.Equal("1.15 GB", t3)

	t4, err4 := NumberToHumanSize(1234567, "precision:2")
	a.NoError(err4)
	a.Equal("1.2 MB", t4)

	t5, err5 := NumberToHumanSize(524288000, "precision:5")
	a.NoError(err5)
	a.Equal("500 MB", t5)

	t6, err6 := NumberToHumanSize(1234567, "separator:,", "precision:2")
	a.NoError(err6)
	a.Equal("1,2 MB", t6)

}
