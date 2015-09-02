package gonumbers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NumberToPhone(val interface{}, args ...interface{}) (string, error) {

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

	_, _, _, delimiter, area_code, extension, country_code, digits_size := func_params(args...)

	if delimiter == "$notset$" {
		delimiter = "-"
	}

	if area_code == nil {
		area_code = false
	}

	if extension == "$notset$" {
		extension = ""
	}

	if country_code == "$notset$" {
		country_code = ""
	}

	if digits_size == nil {
		digits_size = 3
	}

	// Format phone
	var formated_phone string

	phone_sliced := split_to_phone_format(parsed_val, digits_size.(int))

	if area_code.(bool) {
		if len(phone_sliced) == 3 {
			ps_area_code := phone_sliced[0]
			phone_sliced = append(phone_sliced[:0], phone_sliced[1:]...)
			phone_rest_joined := strings.Join(phone_sliced, delimiter)

			formated_phone = fmt.Sprintf("(%s) %s", ps_area_code, phone_rest_joined)
			formated_phone = add_country_code(formated_phone, country_code)

		} else {
			formated_phone = strings.Join(phone_sliced, delimiter)
			formated_phone = add_country_code(formated_phone, country_code)
		}
	} else {

		if country_code != "" {
			new_phone_slice := make([]string, 0)
			old_phone_sliced := phone_sliced
			new_phone_slice = append(new_phone_slice, country_code)
			new_phone_slice = append(new_phone_slice, old_phone_sliced...)
			formated_phone = strings.Join(new_phone_slice, delimiter)
			formated_phone = fmt.Sprintf("+%s", formated_phone)
		} else {
			formated_phone = strings.Join(phone_sliced, delimiter)
		}

	}

	if extension != "" {
		formated_phone = fmt.Sprintf("%s x %s", formated_phone, extension)
	}

	return formated_phone, nil

}

func split_to_phone_format(val string, second_number_len int) []string {

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
