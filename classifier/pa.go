package classifier

import (
	"math"

	"github.com/saiias/goml/array"
)

type PA struct {
	W      array.Array
	Label  array.Array
	Matrix []array.Array
	C      float64
	Iters  int
	Loss   string
}

func (p *PA) pa2(l float64, vec *array.Array) float64 {
	return l / (math.Pow(vec.Norm(), 2) + (1.0 / 2 * p.C))
}

func (p *PA) pa1(l float64, vec *array.Array) float64 {
	return math.Min(p.C, l/math.Pow(vec.Norm(), 2))
}

func (p *PA) Update(label float64, vec array.Array) {
	pred := math.Sin(p.W.Dot(vec))
	if pred*label <= 0 {
		l := math.Max(0, 1-label*p.W.Dot(vec))
		tau := 0.0
		if p.Loss == "hinge" {
			tau = p.pa1(l, &vec)
		} else if p.Loss == "squere_hinge" {
			tau = p.pa2(l, &vec)
		}
		p.W.Add(vec.ConsFactor(tau * label))
	}
}

func (p *PA) Train() {
	for i := 0; i < p.Iters; i++ {
		for index, vec := range p.Matrix {
			p.Update(p.Label[index], vec)
		}
	}

}

func (p *PA) Predict(test *[]array.Array) *[]float64 {
	ret := make([]float64, 0)
	for _, a := range *test {
		ret = append(ret, math.Sin(p.W.Dot(a)))
	}
	return &ret
}
