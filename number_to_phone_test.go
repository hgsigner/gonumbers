package gonumbers_test

import (
	"fmt"
	"testing"

	"github.com/hgsigner/gonumbers"
)

func Test_ToPhone(t *testing.T) {
	// a := assert.New(t)

	// ntph_err1 := gonumbers.ToPhone()
	// ntph_err1_resp, err1 := ntph_err1.Perform("good")
	// a.Error(err1)
	// a.Contains(err1.Error(), "The value should an integer.")
	// a.Equal("", ntph_err1_resp)

	// ntph_err2 := gonumbers.ToPhone()
	// ntph_err2_resp, err2 := ntph_err2.Perform("123abc456")
	// a.Error(err2)
	// a.Contains(err2.Error(), "The value should an integer.")
	// a.Equal("", ntph_err2_resp)

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

	for _, test := range tests {
		ntph := gonumbers.ToPhone()

		if test.addDelimiter {
			ntph.Options(gonumbers.Delimiter(test.delimiter))
		}

		if test.addAreaCode {
			ntph.Options(gonumbers.AreaCode(test.areaCode))
		}

		if test.addExtension {
			ntph.Options(gonumbers.Extension(test.extension))
		}

		if test.addCountryCode {
			ntph.Options(gonumbers.CountryCode(test.countryCode))
		}

		if test.addDigitsSize {
			ntph.Options(gonumbers.DigitsSize(test.digitsSize))
		}

		ntph_final, _ := ntph.Perform(test.in)
		// if ntph_final != test.out {
		// 	t.Errorf("\nExpected: %s\nGot:      %s", test.out, ntph_final)
		// }

		assert(t, test.out, ntph_final)
	}
}

func ExampleToPhone() {
	ntph1 := gonumbers.ToPhone()
	resp1, err1 := ntph1.Perform(5551234)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(resp1)

	ntph2 := gonumbers.ToPhone()
	resp2, err2 := ntph2.Perform(1235551234)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(resp2)

	ntph3 := gonumbers.ToPhone()
	ntph3.Options(gonumbers.AreaCode(true))
	resp3, err3 := ntph3.Perform(1235551234)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(resp3)

	ntph4 := gonumbers.ToPhone()
	ntph4.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("1"))
	resp4, err4 := ntph4.Perform(1235551234)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(resp4)

	ntph5 := gonumbers.ToPhone()
	ntph5.Options(gonumbers.CountryCode("1"))
	resp5, err5 := ntph5.Perform(1235551234)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(resp5)

	ntph6 := gonumbers.ToPhone()
	ntph6.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("1"), gonumbers.Extension("4545"))
	resp6, err6 := ntph6.Perform(1235551234)
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println(resp6)

	ntph7 := gonumbers.ToPhone()
	ntph7.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("1"), gonumbers.Extension("4545"), gonumbers.DigitsSize(5))
	resp7, err7 := ntph7.Perform(1234555556789)
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(resp7)

	ntph8 := gonumbers.ToPhone()
	ntph8.Options(gonumbers.AreaCode(true), gonumbers.CountryCode("55"), gonumbers.Delimiter(","), gonumbers.DigitsSize(5))
	resp8, err8 := ntph8.Perform(1234555556789)
	if err8 != nil {
		fmt.Println(err8)
	}
	fmt.Println(resp8)

	// Output:
	// 555-1234
	// 123-555-1234
	// (123) 555-1234
	// +1(123) 555-1234
	// +1-123-555-1234
	// +1(123) 555-1234 x 4545
	// +1(1234) 55555-6789 x 4545
	// +55(1234) 55555,6789

}
