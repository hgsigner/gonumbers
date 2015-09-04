package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberToDelimiter struct {
	separator, delimiter           string
	isSeparatorSet, isDelimiterSet bool
}

type option func(*NumberToDelimiter)

func (ntd *NumberToDelimiter) Options(options ...option) {
	for _, opt := range options {
		opt(ntd)
	}
}

func (ntd *NumberToDelimiter) Separator(s string) option {
	return func(ntd *NumberToDelimiter) {
		ntd.separator = s
		ntd.isSeparatorSet = true
	}
}

func (ntd *NumberToDelimiter) Delimiter(d string) option {
	return func(ntd *NumberToDelimiter) {
		ntd.delimiter = d
		ntd.isDelimiterSet = true
	}
}

func (ntd *NumberToDelimiter) Perform(n float64) string {

	if !ntd.isSeparatorSet {
		ntd.separator = "."
	}

	if !ntd.isDelimiterSet {
		ntd.delimiter = ","
	}

	s := strings.Split(strconv.FormatFloat(n, 'f', -1, 64), ".")

	gt := group_in_thousands(s[0])
	fp_final := strings.Join(gt, ntd.delimiter)

	var final_value string
	if len(s) > 1 {
		final_value = fmt.Sprintf("%s%s%s", fp_final, ntd.separator, s[1])
	} else {
		final_value = fp_final
	}

	return final_value
}
