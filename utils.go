package gonumbers

import "math"

func group_in_thousands(val string, precision int) []string {
	slice := make([]string, 0)
	times := math.Ceil(float64(len(val)) / float64(precision))

	for i := 0; i <= int(times)-1; i++ {
		from := 0 + (precision * i)
		to := precision + (precision * i)
		if to > len(val) {
			to = len(val)
		}
		slice = append(slice, val[len(val)-to:len(val)-from])
	}

	// Reverse slice

	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}

	return slice
}
