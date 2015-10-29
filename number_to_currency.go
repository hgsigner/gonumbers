package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberToCurrency struct {
	precision                  int
	unit, separator, delimiter string
}

//It sets up options for NumberToCurrency.
//It receives: gonumbers.Precision(int), gonumbers.Unit(string), gonumbers.Separator(string) and gonumbers.Delimiter(string)
func (ntc *NumberToCurrency) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case precisionOption:
			opt.(precisionOption)(ntc)
		case unitOption:
			opt.(unitOption)(ntc)
		case separatorOption:
			opt.(separatorOption)(ntc)
		case delimiterOption:
			opt.(delimiterOption)(ntc)
		}
	}
}

// Implements interfaces
func (ntc *NumberToCurrency) setUnit(s string) {
	ntc.unit = s
}

func (ntc *NumberToCurrency) setPrecision(s int) {
	ntc.precision = s
}

func (ntc *NumberToCurrency) setSeparator(s string) {
	ntc.separator = s
}

func (ntc *NumberToCurrency) setDelimiter(d string) {
	ntc.delimiter = d
}

func (ntc *NumberToCurrency) Perform(n float64) string {

	precision_pattern := fmt.Sprintf(".%df", ntc.precision)
	number_with_precision := fmt.Sprintf("%"+precision_pattern, n)

	parsed_value, _ := strconv.ParseFloat(number_with_precision, 64)

	s := strings.Split(strconv.FormatFloat(parsed_value, 'f', ntc.precision, 64), ".")

	gt := groupInThousands(s[0])
	fp_final := strings.Join(gt, ntc.delimiter)

	final_value := fmt.Sprintf("%s%s%s%s", ntc.unit, fp_final, ntc.separator, s[1])

	return final_value

}
