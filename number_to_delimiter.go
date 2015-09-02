package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberToDelimiter(n float64, args ...interface{}) string {

	// defaults

	_, separator, _, delimiter, _, _, _ := func_params(args...)

	if separator == "$notset$" {
		separator = "."
	}

	if delimiter == "$notset$" {
		delimiter = ","
	}

	s := strings.Split(strconv.FormatFloat(n, 'f', -1, 64), ".")

	gt := group_in_thousands(s[0])
	fp_final := strings.Join(gt, delimiter)

	var final_value string
	if len(s) > 1 {
		final_value = fmt.Sprintf("%s%s%s", fp_final, separator, s[1])
	} else {
		final_value = fp_final
	}

	return final_value

}
