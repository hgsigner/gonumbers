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

func digits_count(n float64) float64 {
	return math.Floor(math.Log10(math.Abs(float64(n))) + 1)
}

func rounded_number(n int64, precision int) string {
	dc := digits_count(float64(n))
	mult := math.Pow10(int(dc) - precision)

	mstring := strconv.FormatFloat(float64(mult), 'f', -1, 64)
	nstring := strconv.FormatFloat(float64(n), 'f', -1, 64)

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
	precision int32,
	delimiter string,
	area_code interface{},
	extension string,
	country_code string,
	digits_size interface{},
) {

	unit = "$notset$"
	separator = "$notset$"
	precision = -1
	delimiter = "$notset$"
	area_code = nil
	extension = "$notset$"
	country_code = "$notset$"
	digits_size = nil

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
		}
	}

	return
}
