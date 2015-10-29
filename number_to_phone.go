package gonumbers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NumberToPhone struct {
	delimiter   string
	areaCode    bool
	extension   string
	countryCode string
	digitsSize  int
}

// Sets up options for NumberToPhone.
// Receives: gonumbers.Delimiter(string),
// gonumbers.AreaCode(bool), gonumbers.Extension(string),
// gonumbers.CountryCode(string) and gonumbers.DigitsSize(int)
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

// Implements interfaces
func (ntph *NumberToPhone) setDelimiter(d string) {
	ntph.delimiter = d
}

func (ntph *NumberToPhone) setAreaCode(ac bool) {
	ntph.areaCode = ac
}

func (ntph *NumberToPhone) setExtension(ex string) {
	ntph.extension = ex
}

func (ntph *NumberToPhone) setCountryCode(cc string) {
	ntph.countryCode = cc
}

func (ntph *NumberToPhone) setDigitsSize(ds int) {
	ntph.digitsSize = ds
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
			return "", notIntegerError
		}
		parsed_val = val.(string)
	case reflect.Int:
		parsed_val = strconv.Itoa(val.(int))
	default:
		return "", notIntegerError
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
			formated_phone = addCountryCode(formated_phone, ntph.countryCode)

		} else {
			formated_phone = strings.Join(phone_sliced, ntph.delimiter)
			formated_phone = addCountryCode(formated_phone, ntph.countryCode)
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

// Helper function that splits an string into a []string
// given a second number length. E.g:
// - val:               "1231235551234",
// - second_number_len: 3,
// - output:            []string{"123123", "555", "1234"}
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

// Adds the area code if necessary.
func addCountryCode(phone, cc string) string {
	if cc != "" {
		return fmt.Sprintf("+%s%s", cc, phone)
	}
	return phone
}
