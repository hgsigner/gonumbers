package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberToHuman struct {
	precision int
	separator string
}

//It sets up options for NumberToHuman.
//It receives: gonumbers.Precision(int) and gonumbers.Separator(string)
func (nth *NumberToHuman) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case separatorOption:
			opt.(separatorOption)(nth)
		case precisionOption:
			opt.(precisionOption)(nth)
		}
	}
}

// Implements interfaces
func (nth *NumberToHuman) setSeparator(s string) {
	nth.separator = s
}

func (nth *NumberToHuman) setPrecision(p int) {
	nth.precision = p
}

//It performs the formating of a float64 number.
func (nth *NumberToHuman) Perform(n float64) string {

	// Inits

	var final_value, gt_first, gt_second string

	val := roundedNumber(n, nth.precision)
	gt := groupInThousands(val)

	if len(gt) > 1 {

		gt_first = gt[0]

		if gt_second = strings.Trim(gt[1], "0"); gt_second == "" {
			gt_second = "0"
		}

		var units string

		switch len(gt) {
		case 1:
			units = ""
		case 2:
			units = "Thousand"
		case 3:
			units = "Million"
		case 4:
			units = "Billion"
		case 5:
			units = "Trillion"
		case 6:
			units = "Quadrillion"
		case 7:
			units = "Quintillion"
		case 8:
			units = "Sextillion"
		case 9:
			units = "Septillion"
		case 10:
			units = "Octillion"
		case 11:
			units = "Nonillion"
		case 12:
			units = "Decillion"
		case 13:
			units = "Undecillion"
		case 14:
			units = "Duodecillion"
		case 15:
			units = "Tredecillion"
		default:
			units = "Googolplex"

		}

		second_to_int, err := strconv.Atoi(gt_second)
		if err != nil {
			return fmt.Sprintf("Can't convert %T into int.", gt_second)
		}

		if second_to_int != 0 {

			switch len(gt_first) {
			case 1:
				final_value = fmt.Sprintf("%s%s%s %s", gt_first, nth.separator, gt_second, units)
			case 2, 3:
				final_value = fmt.Sprintf("%s%s%s %s", gt_first, nth.separator, gt_second[:1], units)
			}

		} else {
			final_value = fmt.Sprintf("%s %s", gt_first, units)
		}

	} else {
		sv := strings.Join(strings.Split(gt[0], "."), nth.separator)
		final_value = sv
	}

	return final_value

}
