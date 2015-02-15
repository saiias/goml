package main

import (
	"fmt"

	"github.com/saiias/goml/array"
	"github.com/saiias/goml/classifier"
	"github.com/saiias/goml/utils"
)

func main() {
	path := "/Users/sai/dataset/a9a"
	label, matrix := utils.Svmlight(path)

	tpath := "/Users/sai/dataset/a9a.t"
	tlabel, tmatrix := utils.Svmlight(tpath)

	pw := array.MakeRandArray(100)
	pe := classifier.Perceptron{
		W:      pw,
		Label:  label,
		Matrix: matrix,
		Eta:    0.01,
		Iters:  1,
	}
	pe.Train()
	peret := pe.Predict(&tmatrix)
	fmt.Printf("Perceptron Acc:%v\n", utils.Accuracy(&tlabel, peret))

	paw := array.MakeArray()
	pa := classifier.PA{
		W:      paw,
		Label:  label,
		Matrix: matrix,
		C:      0.1,
		Iters:  1,
		Loss:   "hinge",
	}

	pa.Train()
	pret := pa.Predict(&tmatrix)
	fmt.Printf("PA-1       Acc:%v\n", utils.Accuracy(&tlabel, pret))

	w := array.MakeArray()
	conv := array.MakeArray()
	ar := classifier.Arow{
		W:          w,
		Label:      label,
		Matrix:     matrix,
		Confidence: conv,
		Iters:      1,
		R:          0.0001,
	}
	ar.Train()
	ret := ar.Predict(&tmatrix)
	fmt.Printf("AROW       Acc:%v\n", utils.Accuracy(&tlabel, ret))

}
