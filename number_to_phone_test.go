package gonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberToPhone(t *testing.T) {
	a := assert.New(t)

	ntph_err1 := NewNumberToPhone()
	ntph_err1_resp, err1 := ntph_err1.Perform("good")
	a.Error(err1)
	a.Contains(err1.Error(), "The value should an integer.")
	a.Equal("", ntph_err1_resp)

	ntph_err2 := NewNumberToPhone()
	ntph_err2_resp, err2 := ntph_err2.Perform("123abc456")
	a.Error(err2)
	a.Contains(err2.Error(), "The value should an integer.")
	a.Equal("", ntph_err2_resp)

	tests := []struct {
		in             interface{}
		out            string
		delimiter      string
		addDelimiter   bool
		areaCode       bool
		addAreaCode    bool
		extension      string
		addExtension   bool
		countryCode    string
		addCountryCode bool
		digitsSize     int
		addDigitsSize  bool
	}{
		{in: 5551234, out: "555-1234"},
		{in: 1235551234, out: "123-555-1234"},
		{
			in:          1235551234,
			addAreaCode: true,
			areaCode:    true,
			out:         "(123) 555-1234",
		},
		{
			in:             1235551234,
			addAreaCode:    true,
			areaCode:       true,
			addCountryCode: true,
			countryCode:    "1",
			out:            "+1(123) 555-1234",
		},
		{
			in:             1235551234,
			addCountryCode: true,
			countryCode:    "1",
			out:            "+1-123-555-1234",
		},
		{
			in:          5551234,
			addAreaCode: true,
			areaCode:    true,
			out:         "555-1234",
		},
		{
			in:             1234555556789,
			addAreaCode:    true,
			areaCode:       true,
			addCountryCode: true,
			countryCode:    "1",
			addExtension:   true,
			extension:      "4545",
			out:            "+1(123455) 555-6789 x 4545",
		},
		{
			in:             1234555556789,
			addAreaCode:    true,
			areaCode:       true,
			addCountryCode: true,
			countryCode:    "1",
			addExtension:   true,
			extension:      "4545",
			addDigitsSize:  true,
			digitsSize:     4,
			out:            "+1(12345) 5555-6789 x 4545",
		},
		{
			in:             1234555556789,
			addAreaCode:    true,
			areaCode:       true,
			addCountryCode: true,
			countryCode:    "1",
			addExtension:   true,
			extension:      "4545",
			addDigitsSize:  true,
			digitsSize:     5,
			out:            "+1(1234) 55555-6789 x 4545",
		},
		{
			in:             1234555556789,
			addDelimiter:   true,
			delimiter:      ",",
			addAreaCode:    true,
			areaCode:       true,
			addCountryCode: true,
			countryCode:    "55",
			addDigitsSize:  true,
			digitsSize:     5,
			out:            "+55(1234) 55555,6789",
		},
	}

	for _, t := range tests {
		ntph := NewNumberToPhone()

		if t.addDelimiter {
			ntph.Options(Delimiter(t.delimiter))
		}

		if t.addAreaCode {
			ntph.Options(AreaCode(t.areaCode))
		}

		if t.addExtension {
			ntph.Options(Extension(t.extension))
		}

		if t.addCountryCode {
			ntph.Options(CountryCode(t.countryCode))
		}

		if t.addDigitsSize {
			ntph.Options(DigitsSize(t.digitsSize))
		}

		ntph_final, _ := ntph.Perform(t.in)
		a.Equal(t.out, ntph_final)
	}
}

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
