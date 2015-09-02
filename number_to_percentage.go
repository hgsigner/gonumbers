package gonumbers

import (
	"fmt"
	"reflect"
	"strconv"
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

	_, _, precision, _ := func_params(args...)

	if precision == -1 {
		precision = 3
	}

	formated_val := strconv.FormatFloat(parsed_float, 'f', int(precision), 64)

	return fmt.Sprintf("%s%%", formated_val)

}
