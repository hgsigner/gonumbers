package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberToCurrency struct {
	precision                                                 int
	unit, separator, delimiter                                string
	isUnitSet, isSeparatorSet, isDelimiterSet, isPrecisionSet bool
}

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

func (ntc *NumberToCurrency) setUnit(s string) {
	ntc.unit = s
	ntc.isUnitSet = true
}

func (ntc *NumberToCurrency) setPrecision(s int) {
	ntc.precision = s
	ntc.isPrecisionSet = true
}

func (ntc *NumberToCurrency) setSeparator(s string) {
	ntc.separator = s
	ntc.isSeparatorSet = true
}

func (ntc *NumberToCurrency) setDelimiter(d string) {
	ntc.delimiter = d
	ntc.isDelimiterSet = true
}

func (ntc *NumberToCurrency) Perform(n float64) string {

	//Default values

	if !ntc.isUnitSet {
		ntc.unit = "$"
	}

	if !ntc.isSeparatorSet {
		ntc.separator = "."
	}

	if !ntc.isPrecisionSet {
		ntc.precision = 2
	}

	if !ntc.isDelimiterSet {
		ntc.delimiter = ","
	}

	precision_pattern := fmt.Sprintf(".%df", ntc.precision)
	number_with_precision := fmt.Sprintf("%"+precision_pattern, n)

	parsed_value, _ := strconv.ParseFloat(number_with_precision, 64)

	s := strings.Split(strconv.FormatFloat(parsed_value, 'f', ntc.precision, 64), ".")

	gt := group_in_thousands(s[0])
	fp_final := strings.Join(gt, ntc.delimiter)

	final_value := fmt.Sprintf("%s%s%s%s", ntc.unit, fp_final, ntc.separator, s[1])

	return final_value

}
