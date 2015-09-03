package gonumbers

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Round number

func round(v float64) float64 {
	if math.Abs(v-math.Ceil(v)) <= 0.5 {
		return math.Ceil(v)
	}
	return math.Floor(v)
}

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

func func_params(args ...interface{}) (
	unit string,
	separator string,
	precision int,
	delimiter string,
	area_code interface{},
	extension string,
	country_code string,
	digits_size interface{},
	prefix string,
) {

	unit = "$notset$"
	separator = "$notset$"
	precision = -1
	delimiter = "$notset$"
	area_code = nil
	extension = "$notset$"
	country_code = "$notset$"
	digits_size = nil
	prefix = "$notset$"

	// Get values

	for _, item := range args {
		opt := strings.SplitN(item.(string), ":", 2)
		switch opt[0] {
		case "unit":
			unit = opt[1]
		case "separator":
			separator = opt[1]
		case "precision":
			_, err := fmt.Sscanf(item.(string), "precision:%d", &precision)
			if err != nil {
				break
			}
		case "delimiter":
			delimiter = opt[1]
		case "area_code":
			b, err := strconv.ParseBool(opt[1])
			if err != nil {
				break
			}
			area_code = b
		case "extension":
			extension = opt[1]
		case "country_code":
			country_code = opt[1]
		case "digits_size":
			ds, err := strconv.Atoi(opt[1])
			if err != nil {
				break
			}
			digits_size = ds
		case "prefix":
			prefix = opt[1]
		}
	}

	return
}
