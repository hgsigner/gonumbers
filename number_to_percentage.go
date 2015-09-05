package gonumbers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NumberToPercentage struct {
	precision                                      int
	separator, delimiter                           string
	isPrecisionSet, isSeparatorSet, isDelimiterSet bool
}

type optionNTP func(*NumberToPercentage)

func (ntp *NumberToPercentage) Options(options ...optionNTP) {
	for _, opt := range options {
		opt(ntp)
	}
}

func (ntp *NumberToPercentage) Precision(p int) optionNTP {
	return func(ntp *NumberToPercentage) {
		ntp.precision = p
		ntp.isPrecisionSet = true
	}
}

func (ntp *NumberToPercentage) Separator(s string) optionNTP {
	return func(ntp *NumberToPercentage) {
		ntp.separator = s
		ntp.isSeparatorSet = true
	}
}

func (ntp *NumberToPercentage) Delimiter(d string) optionNTP {
	return func(ntp *NumberToPercentage) {
		ntp.delimiter = d
		ntp.isDelimiterSet = true
	}
}

//It performs the convertion of the input to percentage.
//It accepts an interface, but only values related to numbers
//(e.g. floats, ints) or a string that could be parsed
//to float (e.g. "12345").
//Inputs that are not related to numbers will be returned in the
//0.000% format.
func (ntp *NumberToPercentage) Perform(val interface{}) string {

	parsed_float := 0.0

	// Parse val to number

	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.String:
		stf, _ := strconv.ParseFloat(val.(string), 64)
		parsed_float = stf
	case reflect.Int:
		parsed_float = float64(val.(int))
	case reflect.Float64:
		parsed_float = val.(float64)
	}

	// Defaults

	if !ntp.isSeparatorSet {
		ntp.separator = "."
	}

	if !ntp.isPrecisionSet {
		ntp.precision = 3
	}

	if !ntp.isDelimiterSet {
		ntp.delimiter = ","
	}

	formated_val_splited := strings.Split(strconv.FormatFloat(parsed_float, 'f', ntp.precision, 64), ".")
	gt := group_in_thousands(formated_val_splited[0])
	joined_thousands := strings.Join(gt, ntp.delimiter)

	var final_value string
	if len(formated_val_splited) > 1 {

		final_value = fmt.Sprintf("%s%s%s%%", joined_thousands, ntp.separator, formated_val_splited[1])
	} else {
		final_value = fmt.Sprintf("%s%%", joined_thousands)
	}

	return final_value

}
