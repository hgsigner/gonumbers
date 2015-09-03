package gonumbers

import (
	"fmt"
	"math"
)

func NumberToHumanSize(n float64, args ...interface{}) (string, error) {

	st := []string{"byte", "kb", "mb", "gb", "tb"}
	base := 1024

	max := len(st) - 1
	exp := int(math.Log(n) / math.Log(float64(base)))

	if exp > max {
		exp = max
	}

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

	human_size := n / math.Pow(float64(base), float64(exp))
	rounded_n := rounded_number(human_size, 3)

	final_formated := fmt.Sprintf("%v %s", rounded_n, exp_name)

	return final_formated, nil

}
