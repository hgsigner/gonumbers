package gonumbers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NumberToPhone struct {
	delimiter        string
	isDelimiterSet   bool
	areaCode         bool
	isAreaCodeSet    bool
	extension        string
	isExtensionSet   bool
	countryCode      string
	isCountryCodeSet bool
	digitsSize       int
	isDigitsSizeSet  bool
}

//It sets up options for NumberToPhone.
//It receives: gonumbers.Delimiter(string), gonumbers.AreaCode(bool), gonumbers.Extension(string), gonumbers.CountryCode(string) and gonumbers.DigitsSize(int)
func (ntph *NumberToPhone) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case delimiterOption:
			opt.(delimiterOption)(ntph)
		case areaCodeOption:
			opt.(areaCodeOption)(ntph)
		case extensionOption:
			opt.(extensionOption)(ntph)
		case countryCodeOption:
			opt.(countryCodeOption)(ntph)
		case digitsSizeOption:
			opt.(digitsSizeOption)(ntph)
		}
	}
}

func (ntph *NumberToPhone) setDelimiter(d string) {
	ntph.delimiter = d
	ntph.isDelimiterSet = true
}

func (ntph *NumberToPhone) setAreaCode(ac bool) {
	ntph.areaCode = ac
	ntph.isAreaCodeSet = true
}

func (ntph *NumberToPhone) setExtension(ex string) {
	ntph.extension = ex
	ntph.isExtensionSet = true
}

func (ntph *NumberToPhone) setCountryCode(cc string) {
	ntph.countryCode = cc
	ntph.isCountryCodeSet = true
}

func (ntph *NumberToPhone) setDigitsSize(ds int) {
	ntph.digitsSize = ds
	ntph.isDigitsSizeSet = true
}

//It performs the convertion of the input into phone number format.
//It accepts an interface, but only values related to numbers
//(e.g. floats, ints) or a string that could be parsed
//to float (e.g. "12345").
//Inputs that are not related to numbers will return an error.
func (ntph *NumberToPhone) Perform(val interface{}) (string, error) {

	// Parse val

	var parsed_val string

	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.String:
		_, err := strconv.ParseInt(val.(string), 0, 64)
		if err != nil {
			return "", errors.New("The value should an integer.")
		}
		parsed_val = val.(string)
	case reflect.Int:
		parsed_val = strconv.Itoa(val.(int))
	default:
		return "", errors.New("The value should an integer.")
	}

	// Defaults

	if !ntph.isDelimiterSet {
		ntph.delimiter = "-"
	}

	if !ntph.isAreaCodeSet {
		ntph.areaCode = false
	}

	if !ntph.isExtensionSet {
		ntph.extension = ""
	}

	if !ntph.isCountryCodeSet {
		ntph.countryCode = ""
	}

	if !ntph.isDigitsSizeSet {
		ntph.digitsSize = 3
	}

	// Format phone
	var formated_phone string

	phone_sliced := splitToPhoneFormart(parsed_val, ntph.digitsSize)

	if ntph.areaCode {
		if len(phone_sliced) == 3 {
			ps_area_code := phone_sliced[0]
			phone_sliced = append(phone_sliced[:0], phone_sliced[1:]...)
			phone_rest_joined := strings.Join(phone_sliced, ntph.delimiter)

			formated_phone = fmt.Sprintf("(%s) %s", ps_area_code, phone_rest_joined)
			formated_phone = add_country_code(formated_phone, ntph.countryCode)

		} else {
			formated_phone = strings.Join(phone_sliced, ntph.delimiter)
			formated_phone = add_country_code(formated_phone, ntph.countryCode)
		}
	} else {

		if ntph.countryCode != "" {
			new_phone_slice := make([]string, 0)
			old_phone_sliced := phone_sliced
			new_phone_slice = append(new_phone_slice, ntph.countryCode)
			new_phone_slice = append(new_phone_slice, old_phone_sliced...)
			formated_phone = strings.Join(new_phone_slice, ntph.delimiter)
			formated_phone = fmt.Sprintf("+%s", formated_phone)
		} else {
			formated_phone = strings.Join(phone_sliced, ntph.delimiter)
		}

	}

	if ntph.extension != "" {
		formated_phone = fmt.Sprintf("%s x %s", formated_phone, ntph.extension)
	}

	return formated_phone, nil
}

func splitToPhoneFormart(val string, second_number_len int) []string {

	final_slice := make([]string, 0)

	phone_last_part := val[len(val)-4 : len(val)]
	scn_index := 4 + second_number_len
	phone_second_part := val[len(val)-scn_index : len(val)-4]
	phone_area_code := val[0 : len(val)-scn_index]

	if phone_area_code != "" {
		final_slice = append(final_slice, phone_area_code)
	}
	final_slice = append(final_slice, phone_second_part)
	final_slice = append(final_slice, phone_last_part)

	return final_slice

}

func add_country_code(phone, cc string) string {

	if cc != "" {
		return fmt.Sprintf("+%s%s", cc, phone)
	}

	return phone

}
