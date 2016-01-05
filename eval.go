package utils

import (
	"github.com/saiias/goml/array"
)

func Accuracy(label *array.Array, pred *[]float64) float64 {
	count := len(*label)
	p := 0.0
	for i, l := range *label {
		if (l > 0 && (*pred)[i] > 0) || (l < 0 && (*pred)[i] < 0) {
			p += 1.0
		}
	}
	return p / float64(count)
}
