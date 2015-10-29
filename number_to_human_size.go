package gonumbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NumberToHumanSize struct {
	precision int
	separator string
	delimiter string
	prefix    string
}

//It sets up options for NumberToHumanSize.
//It receives: gonumbers.Precision(int), gonumbers.Separator(string) and gonumbers.Delimiter(string) and gonumbers.Prefix(string)
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

// Implements interfaces.
func (nths *NumberToHumanSize) setPrecision(p int) {
	nths.precision = p
}

func (nths *NumberToHumanSize) setSeparator(s string) {
	nths.separator = s
}

func (nths *NumberToHumanSize) setDelimiter(d string) {
	nths.delimiter = d
}

func (nths *NumberToHumanSize) setPrefix(p string) {
	nths.prefix = p
}

// Performs the formating of a float64 number.
// If the prefix is not binary or si, it returns an error.
func (nths *NumberToHumanSize) Perform(n float64) (string, error) {

	st := []string{"byte", "kb", "mb", "gb", "tb"}
	var base float64

	switch nths.prefix {
	case "binary":
		base = 1024
	case "si":
		base = 1000
	default:
		return "", unknownBinaryPrefix
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
	rounded_n1 := roundedNumber(human_size, nths.precision)
	prounded_n1, _ := strconv.ParseFloat(rounded_n1, 64)
	rounded_n2 := roundWithPrecision(prounded_n1, nths.precision)

	final_round_splited := strings.Split(strconv.FormatFloat(rounded_n2, 'f', -1, 64), ".")

	final_round_joined := strings.Join(final_round_splited, nths.separator)

	final_formated := fmt.Sprintf("%v %s", final_round_joined, exp_name)

	return final_formated, nil

}
