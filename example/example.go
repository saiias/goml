package main

import (
	"fmt"
	"github.com/saiias/goml"
)

func main(){
	path := "/Users/sai/dataset/a9a"
	label,matrix := goml.Svmlight(path)

	tpath := "/Users/sai/dataset/a9a.t"
	tlabel,tmatrix := goml.Svmlight(tpath)

	pw := goml.MakeRandArray(100)
	pe := goml.Perceptron{
		W: pw,
		Label : label,
		Matrix: matrix,
		Eta: 0.001,
		Iters:1,
	}
	pe.Train()
	peret := pe.Predict(&tmatrix)
	fmt.Printf("Perceptron Acc:%v\n",goml.Accuracy(&tlabel,peret))

	paw := goml.MakeArray()
	pa := goml.PA{
		W:paw,
		Label:label,
		Matrix: matrix,
		C:0.1,
		Iters:1,
		Loss:"hinge",
	}

	pa.Train()
	pret := pa.Predict(&tmatrix)
	fmt.Printf("PA-1       Acc:%v\n",goml.Accuracy(&tlabel,pret))

	w := goml.MakeArray()
	conv := goml.MakeArray()
	ar := goml.Arow{
		W: w,
		Label: label,
		Matrix: matrix,
		Confidence:conv,
		Iters: 1,
		R:0.0001,
	}
	ar.Train()
	ret := ar.Predict(&tmatrix)
	fmt.Printf("AROW       Acc:%v\n",goml.Accuracy(&tlabel,ret))


}
