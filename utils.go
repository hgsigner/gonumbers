package gonumbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func group_in_thousands(val string, precision int) []string {
	slice := make([]string, 0)
	times := math.Ceil(float64(len(val)) / float64(precision))

	if len(val) >= precision {

		for i := 0; i <= int(times)-1; i++ {
			from := 0 + (precision * i)
			to := precision + (precision * i)
			if to >= len(val) {
				to = len(val)
			}
			slice = append(slice, val[len(val)-to:len(val)-from])
		}

		// Reverse slice

		for i := len(slice)/2 - 1; i >= 0; i-- {
			opp := len(slice) - 1 - i
			slice[i], slice[opp] = slice[opp], slice[i]
		}

	} else {
		expv := math.Pow10(precision)
		pval, _ := strconv.Atoi(val)
		res := float64(pval) / expv
		presult := strconv.FormatFloat(res, 'f', precision, 64)
		slice = append(slice, presult)
	}

	return slice
}

func func_params(args ...interface{}) (
	separator string,
	precision int,
) {

	// Defaults

	separator = "."
	precision = 3

	// Get values

	for _, item := range args {
		opt := strings.SplitN(item.(string), ":", 2)
		switch opt[0] {
		case "separator":
			_, err := fmt.Sscanf(item.(string), "separator:%s", &separator)
			if err != nil {
				break
			}
		case "precision":
			_, err := fmt.Sscanf(item.(string), "precision:%d", &precision)
			if err != nil {
				break
			}
		}
	}

	return
}
