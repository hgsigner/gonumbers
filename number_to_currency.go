package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberToCurrency(n float64, args ...interface{}) string {

	//Default values

	unit, separator, precision, delimiter, _, _, _, _ := func_params(args...)

	if unit == "$notset$" {
		unit = "$"
	}

	if separator == "$notset$" {
		separator = "."
	}

	if precision == -1 {
		precision = 2
	}

	if delimiter == "$notset$" {
		delimiter = ","
	}

	precision_pattern := fmt.Sprintf(".%df", precision)
	number_with_precision := fmt.Sprintf("%"+precision_pattern, float64(n))

	parsed_value, _ := strconv.ParseFloat(number_with_precision, 64)

	s := strings.Split(strconv.FormatFloat(parsed_value, 'f', int(precision), 64), ".")

	gt := group_in_thousands(s[0])
	fp_final := strings.Join(gt, delimiter)

	final_value := fmt.Sprintf("%s%s%s%s", unit, fp_final, separator, s[1])

	return final_value
}
