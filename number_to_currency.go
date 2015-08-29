package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberToCurrency(n float64, args ...interface{}) string {

	//Default values

	precision := 2
	unit := "$"
	separator := "."
	delimiter := ","

	switch len(args) {
	case 1:
		precision = args[0].(int)
	case 2:
		precision = args[0].(int)
		unit = args[1].(string)
	case 3:
		precision = args[0].(int)
		unit = args[1].(string)
		separator = args[2].(string)
	case 4:
		precision = args[0].(int)
		unit = args[1].(string)
		separator = args[2].(string)
		delimiter = args[3].(string)
	}

	precision_pattern := fmt.Sprintf(".%df", precision)
	number_with_precision := fmt.Sprintf("%"+precision_pattern, float64(n))

	parsed_value, err := strconv.ParseFloat(number_with_precision, 64)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Split(strconv.FormatFloat(parsed_value, 'f', precision, 64), ".")

	gt := group_in_thousands(s[0], 3)
	fp_final := strings.Join(gt, delimiter)

	final_value := fmt.Sprintf("%s%s%s%s", unit, fp_final, separator, s[1])

	return final_value
}
