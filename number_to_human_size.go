package gonumbers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NumberToHumanSize struct {
	precision      int
	separator      string
	delimiter      string
	prefix         string
	isSeparatorSet bool
	isDelimiterSet bool
	isPrecisionSet bool
	isPrefixSet    bool
}

func (nths *NumberToHumanSize) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case precisionOption:
			opt.(precisionOption)(nths)
		case separatorOption:
			opt.(separatorOption)(nths)
		case delimiterOption:
			opt.(delimiterOption)(nths)
		case prefixOption:
			opt.(prefixOption)(nths)
		}
	}
}

func (nths *NumberToHumanSize) setPrecision(p int) {
	nths.precision = p
	nths.isPrecisionSet = true
}

func (nths *NumberToHumanSize) setSeparator(s string) {
	nths.separator = s
	nths.isSeparatorSet = true
}

func (nths *NumberToHumanSize) setDelimiter(d string) {
	nths.delimiter = d
	nths.isDelimiterSet = true
}

func (nths *NumberToHumanSize) setPrefix(p string) {
	nths.prefix = p
	nths.isPrefixSet = true
}

func (nths *NumberToHumanSize) Perform(n float64) (string, error) {

	st := []string{"byte", "kb", "mb", "gb", "tb"}
	var base float64

	// Defaults

	if !nths.isSeparatorSet {
		nths.separator = "."
	}

	if !nths.isPrecisionSet {
		nths.precision = 3
	}

	if !nths.isDelimiterSet {
		nths.delimiter = ""
	}

	if !nths.isPrefixSet {
		nths.prefix = "binary"
	}

	switch nths.prefix {
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
	rounded_n1 := rounded_number(human_size, nths.precision)
	prounded_n1, _ := strconv.ParseFloat(rounded_n1, 64)
	rounded_n2 := round_with_precision(prounded_n1, nths.precision)

	final_round_splited := strings.Split(strconv.FormatFloat(rounded_n2, 'f', -1, 64), ".")

	final_round_joined := strings.Join(final_round_splited, nths.separator)

	final_formated := fmt.Sprintf("%v %s", final_round_joined, exp_name)

	return final_formated, nil

}
