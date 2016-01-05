package array

import (
	"testing"
)

func TestDot(t *testing.T) {
	v1 := Array{
		1: 1.0,
		2: 0.5,
		3: 1.0,
	}

	v2 := Array{
		1: 0.5,
		2: 1.0,
		3: 1.0,
	}

	actual := v1.Dot(v2)
	want := 2.0
	if actual != want {
		t.Errorf("want %v,but actual %v", want, actual)
	}

}

func TestUpdate(t *testing.T) {
	v1 := Array{
		1: 1.0,
		2: 0.5,
		3: 1.0,
	}

	v1.Update(1, 0.5)

	actual := v1[1]
	want := 0.5
	if actual != want {
		t.Errorf("want %v,but actual %v", want, actual)
	}

}

func BenchmarkDot(b *testing.B) {
	v1 := Array{
		1: 1.0,
		2: 0.5,
		3: 1.0,
	}

	v2 := Array{
		1: 0.5,
		2: 1.0,
		3: 1.0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1.Dot(v2)
	}
}
