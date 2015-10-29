package gonumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type NumberToDelimiter struct {
	separator, delimiter string
}

//It sets up options for NumberToDelimiter.
//It receives: gonumbers.Separator(string) and gonumbers.Delimiter(string)
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
}

func (ntd *NumberToDelimiter) setDelimiter(d string) {
	ntd.delimiter = d
}

//It performs the formating of a float64 number.
func (ntd *NumberToDelimiter) Perform(n float64) string {

	s := strings.Split(strconv.FormatFloat(n, 'f', -1, 64), ".")

	gt := groupInThousands(s[0])
	fp_final := strings.Join(gt, ntd.delimiter)

	var final_value string
	if len(s) > 1 {
		final_value = fmt.Sprintf("%s%s%s", fp_final, ntd.separator, s[1])
	} else {
		final_value = fp_final
	}

	return final_value
}
