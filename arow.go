package goml

type Arow struct {
	W          Array
	Label      Array
	Matrix     []Array
	Confidence Array
	R          float64
	Iters      int
}

func (a *Arow) Train() {
	for i := 0; i < a.Iters; i++ {
		for index, vec := range a.Matrix {
			a.Update(a.Label[index], vec)
		}
	}
}

func (a *Arow) Update(label float64, vec Array) {
	margin := a.W.Dot(vec)

	if margin*label >= 1.0 {
		return
	}

	confidence := 0.0
	for i, value := range vec {
		if _, ok := a.Confidence[i]; ok {
			confidence += a.Confidence[i] * value * value
		} else {
			confidence += value * value
		}
	}

	beta := 1.0 / (confidence + a.R)
	alpha := (1.0 - label*margin) * beta

	for i, value := range vec {

		if _, ok := a.Confidence[i]; !ok {
			a.Confidence[i] = 1.0
		}

		a.W[i] += alpha * a.Confidence[i] * value * label
		a.Confidence[i] = 1.0 / ((1.0 / a.Confidence[i]) + value*value/a.R)
	}
}

func (a *Arow) Predict(test *[]Array) *[]float64 {
	ret := make([]float64, 0)
	for _, b := range *test {
		ret = append(ret, a.W.Dot(b))
	}
	return &ret
}
