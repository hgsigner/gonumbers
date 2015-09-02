package gonumbers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NumberToPercentage(val interface{}, args ...interface{}) string {

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

	_, separator, precision, delimiter := func_params(args...)

	if separator == "$notset$" {
		separator = "."
	}

	if precision == -1 {
		precision = 3
	}

	if delimiter == "$notset$" {
		delimiter = ","
	}

	formated_val_splited := strings.Split(strconv.FormatFloat(parsed_float, 'f', int(precision), 64), ".")
	gt := group_in_thousands(formated_val_splited[0])
	joined_thousands := strings.Join(gt, delimiter)

	var final_value string
	if len(formated_val_splited) > 1 {

		final_value = fmt.Sprintf("%s%s%s%%", joined_thousands, separator, formated_val_splited[1])
	} else {
		final_value = fmt.Sprintf("%s%%", joined_thousands)
	}

	return final_value

}
