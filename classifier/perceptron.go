package classifier

import (
	"math"

	"github.com/saiias/goml/array"
)

type Perceptron struct {
	W      array.Array
	Label  array.Array
	Matrix []array.Array
	Eta    float64
	Iters  int
}

func (p *Perceptron) Update(label float64, vec array.Array) {
	pred := math.Sin(p.W.Dot(vec))
	if pred*label <= 0 {
		grad := vec.ConsFactor(p.Eta * pred)
		p.W.Add(grad)
	}
}

func (p *Perceptron) Train() {
	for i := 0; i < p.Iters; i++ {
		for index, vec := range p.Matrix {
			p.Update(p.Label[index], vec)
		}
	}
}

func (p *Perceptron) Predict(test *[]array.Array) *[]float64 {
	ret := make([]float64, 0)
	for _, a := range *test {
		ret = append(ret, math.Sin(p.W.Dot(a)))
	}
	return &ret
}
