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

type optionNTC func(*NumberToCurrency)

func (ntc *NumberToCurrency) Options(options ...optionNTC) {
	for _, opt := range options {
		opt(ntc)
	}
}

func (ntc *NumberToCurrency) Unit(s string) optionNTC {
	return func(ntc *NumberToCurrency) {
		ntc.unit = s
		ntc.isUnitSet = true
	}
}

func (ntc *NumberToCurrency) Precision(s int) optionNTC {
	return func(ntc *NumberToCurrency) {
		ntc.precision = s
		ntc.isPrecisionSet = true
	}
}

func (ntc *NumberToCurrency) Separator(s string) optionNTC {
	return func(ntc *NumberToCurrency) {
		ntc.separator = s
		ntc.isSeparatorSet = true
	}
}

func (ntc *NumberToCurrency) Delimiter(d string) optionNTC {
	return func(ntc *NumberToCurrency) {
		ntc.delimiter = d
		ntc.isDelimiterSet = true
	}
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
