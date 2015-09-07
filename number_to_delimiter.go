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

func (ntd *NumberToDelimiter) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case separatorOption:
			opt.(separatorOption)(ntd)
		case delimiterOption:
			opt.(delimiterOption)(ntd)
		}
	}
}

func (ntd *NumberToDelimiter) setSeparator(s string) {
	ntd.separator = s
	ntd.isSeparatorSet = true
}

func (ntd *NumberToDelimiter) setDelimiter(d string) {
	ntd.delimiter = d
	ntd.isDelimiterSet = true
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
