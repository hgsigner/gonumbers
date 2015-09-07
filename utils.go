package gonumbers

import (
	"math"
	"math/big"
	"strconv"
)

// Rounds number
func round(v float64) float64 {
	if math.Abs(v-math.Ceil(v)) <= 0.5 {
		return math.Ceil(v)
	}
	return math.Floor(v)
}

// Rounds floats (e.g. 1.200000000001 -> 1.2) to a givan precision
func round_with_precision(val float64, precision int) float64 {
	var round float64
	pow := math.Pow(10, float64(precision))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func digits_count(n float64) int {
	dc := math.Floor(math.Log10(math.Abs(n)) + 1)
	return int(dc)
}

// Rounds rational numbers to a given precision
func rounded_number(n float64, precision int) string {
	dc := digits_count(n)
	mult := math.Pow10(dc - precision)

	mstring := strconv.FormatFloat(float64(mult), 'f', -1, 64)
	nstring := strconv.FormatFloat(n, 'f', -1, 64)

	mb, nb := new(big.Rat), new(big.Rat)
	nb.SetString(nstring)
	mb.SetString(mstring)

	brb, _ := new(big.Rat).Quo(nb, mb).Float64()
	rb := round(brb) * mult

	return strconv.FormatFloat(rb, 'f', -1, 64)
}

//Splits a string of numbers into groups of thousands.
//E.g. "123456" => []string{"123", "456"}
func group_in_thousands(val string) []string {
	slice := make([]string, 0)
	times := math.Ceil(float64(len(val)) / 3.0)

	if len(val) >= 3 {

		for i := 0; i <= int(times)-1; i++ {
			from := 0 + (3 * i)
			to := 3 + (3 * i)
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
		// expv := math.Pow10(precision)
		// pval, _ := strconv.Atoi(val)
		// res := float64(pval) / expv
		// presult := strconv.FormatFloat(res, 'f', precision, 64)
		// slice = append(slice, presult)

		slice = append(slice, val)
	}

	return slice
}
