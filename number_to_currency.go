package gonumbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func NumberToCurrency(n float64, precision int, unit, separator, delimiter string) string {

	precision_pattern := fmt.Sprintf(".%df", precision)
	number_with_precision := fmt.Sprintf("%"+precision_pattern, float64(n))

	parsed_value, err := strconv.ParseFloat(number_with_precision, 64)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Split(strconv.FormatFloat(parsed_value, 'f', precision, 64), ".")

	//It formats the first port of the number

	fpf := make([]string, 0)
	times := math.Ceil(float64(len(s[0])) / 3.0)

	for i := 0; i <= int(times)-1; i++ {
		from := 0 + (3 * i)
		to := 3 + (3 * i)
		if to > len(s[0]) {
			to = len(s[0])
		}
		fpf = append(fpf, s[0][len(s[0])-to:len(s[0])-from])
	}

	// Reverse slice

	for i := len(fpf)/2 - 1; i >= 0; i-- {
		opp := len(fpf) - 1 - i
		fpf[i], fpf[opp] = fpf[opp], fpf[i]
	}

	fp_final := strings.Join(fpf, delimiter)

	final_value := fmt.Sprintf("%s%s%s%s", unit, fp_final, separator, s[1])

	return final_value
}
