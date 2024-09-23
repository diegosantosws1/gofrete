package internal

import (
	"strconv"
	"strings"
)

func ConvertStringToFloa64(str string) (value float64, err error) {
	v, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return
	}
	value = (v / 100)

	return
}
