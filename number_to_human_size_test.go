package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToHumanSize(t *testing.T) {
	a := assert.New(t)

	nths1 := new(NumberToHumanSize)
	nths1.Options(nths1.Prefix("ahhaha"))
	nths1_final, err1 := nths1.Perform(123)
	a.Error(err1)
	a.Contains(err1.Error(), "Prefix must be binary or si.")
	a.Equal("", nths1_final)

	nths2 := new(NumberToHumanSize)
	nths2_final, err2 := nths2.Perform(333)
	a.NoError(err2)
	a.Equal("333 Bytes", nths2_final)

	nths3 := new(NumberToHumanSize)
	nths3_final, err3 := nths3.Perform(1234)
	a.NoError(err3)
	a.Equal("1.21 KB", nths3_final)

	nths4 := new(NumberToHumanSize)
	nths4_final, err4 := nths4.Perform(1234567890)
	a.NoError(err4)
	a.Equal("1.15 GB", nths4_final)

	nths5 := new(NumberToHumanSize)
	nths5.Options(nths5.Precision(2))
	nths5_final, err5 := nths5.Perform(1234567)
	a.NoError(err5)
	a.Equal("1.2 MB", nths5_final)

	nths6 := new(NumberToHumanSize)
	nths6.Options(nths5.Precision(5))
	nths6_final, err6 := nths6.Perform(524288000)
	a.NoError(err6)
	a.Equal("500 MB", nths6_final)

	nths7 := new(NumberToHumanSize)
	nths7.Options(nths7.Precision(2), nths7.Separator(","))
	nths7_final, err7 := nths7.Perform(1234567)
	a.NoError(err7)
	a.Equal("1,2 MB", nths7_final)

}
