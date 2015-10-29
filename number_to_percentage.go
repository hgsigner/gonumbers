package gonumbers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NumberToPercentage struct {
	precision            int
	separator, delimiter string
}

//It sets up options for NumberToPercentage.
//It receives: gonumbers.Precision(int), gonumbers.Separator(string) and gonumbers.Delimiter(string)
func (ntp *NumberToPercentage) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case precisionOption:
			opt.(precisionOption)(ntp)
		case separatorOption:
			opt.(separatorOption)(ntp)
		case delimiterOption:
			opt.(delimiterOption)(ntp)
		}
	}
}

func (ntp *NumberToPercentage) setPrecision(p int) {
	ntp.precision = p
}

func (ntp *NumberToPercentage) setSeparator(s string) {
	ntp.separator = s
}

func (ntp *NumberToPercentage) setDelimiter(d string) {
	ntp.delimiter = d
}

//It performs the convertion of the input into percentage.
//It accepts an interface, but only values related to numbers
//(e.g. floats, ints) or a string that could be parsed
//to float (e.g. "12345").
//Inputs that are not related to numbers will be returned in the
//0.000% format.
func (ntp *NumberToPercentage) Perform(val interface{}) string {

	parsed_float := 0.0

	// Parse val to number

	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.String:
		stf, _ := strconv.ParseFloat(val.(string), 64)
		parsed_float = stf
	case reflect.Int:
		parsed_float = float64(val.(int))
	case reflect.Float64:
		parsed_float = val.(float64)
	}

	formated_val_splited := strings.Split(strconv.FormatFloat(parsed_float, 'f', ntp.precision, 64), ".")
	gt := groupInThousands(formated_val_splited[0])
	joined_thousands := strings.Join(gt, ntp.delimiter)

	var final_value string
	if len(formated_val_splited) > 1 {

		final_value = fmt.Sprintf("%s%s%s%%", joined_thousands, ntp.separator, formated_val_splited[1])
	} else {
		final_value = fmt.Sprintf("%s%%", joined_thousands)
	}

	return final_value

}
