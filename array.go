package goml

import (
	"math"
	"math/rand"
	"sync"
)

var mutex = &sync.Mutex{}

type Array map[int]float64

func (a Array) Dot(b Array) float64 {
	s := 0.0
	for i, elem := range a {
		if value, ok := b[i]; ok {
			s += elem * value
		}
	}
	return s
}

func MakeArray() Array {
	return make(Array, 0)
}

func MakeRandArray(n int) Array {
	vec := make(Array, 0)
	for i := 0; i < n; i++ {
		vec[i] = rand.Float64()
	}
	return vec
}

func Trans(a *map[int]float64) Array {
	return *a
}

func (a *Array) ConsFactor(f float64) *Array {
	mutex.Lock()
	defer mutex.Unlock()
	vec := make(Array, 0)
	for i, elem := range *a {
		vec[i] = f * elem
	}
	return &vec
}

func (a *Array) Add(b *Array) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, elem := range *b {
		if _, ok := (*a)[i]; ok {
			(*a)[i] += elem
		} else {
			(*a)[i] = elem
		}
	}

}

func (a *Array) Update(i int, v float64) {
	mutex.Lock()
	defer mutex.Unlock()
	(*a)[i] = v
}

func (a *Array) Norm() float64 {
	sum := 0.0
	for _, v := range *a {
		sum += math.Pow(v, 2)
	}
	return math.Sqrt(sum)
}

func (a Array) Len() int {
	return len(a)
}

func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Array) Less(i, j int) bool {
	return i < j
}
