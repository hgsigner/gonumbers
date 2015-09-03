package gonumbers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func NumberToHumanSize(n float64, args ...interface{}) (string, error) {

	st := []string{"byte", "kb", "mb", "gb", "tb"}
	var base float64

	// Defaults

	_, separator, precision, delimiter, _, _, _, _, prefix := func_params(args...)

	if separator == "$notset$" {
		separator = "."
	}

	if precision == -1 {
		precision = 3
	}

	if delimiter == "$notset$" {
		delimiter = ""
	}

	if prefix == "$notset$" {
		prefix = "binary"
	}

	switch prefix {
	case "binary":
		base = 1024
	case "si":
		base = 1000
	default:
		return "", errors.New("Prefix must be binary or si.")
	}

	//Calc index of exponent

	max := len(st) - 1
	exp := int(math.Log(n) / math.Log(base))

	if exp > max {
		exp = max
	}

	// Retrives the exponent title
	var exp_name string
	switch st[exp] {
	case "byte":
		exp_name = "Bytes"
	case "kb":
		exp_name = "KB"
	case "mb":
		exp_name = "MB"
	case "gb":
		exp_name = "GB"
	case "tr":
		exp_name = "TR"
	}

	// Retrieves number and formats it

	human_size := n / math.Pow(base, float64(exp))
	rounded_n1 := rounded_number(human_size, precision)
	prounded_n1, _ := strconv.ParseFloat(rounded_n1, 64)
	rounded_n2 := round_with_precision(prounded_n1, precision)

	final_round_splited := strings.Split(strconv.FormatFloat(rounded_n2, 'f', -1, 64), ".")

	final_round_joined := strings.Join(final_round_splited, separator)

	final_formated := fmt.Sprintf("%v %s", final_round_joined, exp_name)

	return final_formated, nil

}
