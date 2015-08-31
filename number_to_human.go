package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberToHuman(num int, args ...interface{}) string {

	//Inits

	var final_value, gt_first, gt_second string

	// defaults

	separator, precision := func_params(args...)

	// convert number

	val := strconv.Itoa(num)
	gt := group_in_thousands(val, precision)

	if len(gt) > 1 {

		gt_first = gt[0]
		gt_second = gt[1]

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
			return fmt.Sprintf("Can't convert %s into int.", gt_second)
		}

		if second_to_int != 0 {

			switch len(gt_first) {
			case 1:
				final_value = fmt.Sprintf("%s%s%s %s", gt_first, separator, gt_second, units)
			case 2, 3:
				final_value = fmt.Sprintf("%s%s%s %s", gt_first, separator, gt_second[:1], units)
			}

		} else {
			final_value = fmt.Sprintf("%s %s", gt_first, units)
		}

	} else {
		sv := strings.Join(strings.Split(gt[0], "."), separator)
		final_value = sv
	}

	return final_value

}
