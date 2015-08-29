package gonumbers

import (
	"fmt"
	"strconv"
)

func NumberToHuman(num int) string {

	// defaults

	separator := "."
	precision := 3

	// convert number

	val := strconv.Itoa(num)
	gt := group_in_thousands(val, precision)

	var units string
	var final_value string

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

	switch len(gt[0]) {
	case 1:
		final_value = fmt.Sprintf("%s%s%s %s", gt[0], separator, gt[1], units)
	case 2, 3:
		final_value = fmt.Sprintf("%s%s%s %s", gt[0], separator, gt[1][:1], units)
	}

	return final_value

}
