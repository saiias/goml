package goml

import (
	"math"
)

type Perceptron struct {
	W      Array
	Label  Array
	Matrix []Array
	Eta    float64
	Iters  int
}

func (p *Perceptron) Update(label float64, vec Array) {
	pred := p.W.Dot(vec)
	if pred*label <= 0 {
		grad := vec.ConsFactor(p.Eta * label * pred)
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

func (p *Perceptron) Predict(test *[]Array) *[]float64 {
	ret := make([]float64, 0)
	for _, a := range *test {
		ret = append(ret, math.Sin(p.W.Dot(a)))
	}
	return &ret
}
