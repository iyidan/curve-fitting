package base

import (
	"testing"
)

func TestPoint(t *testing.T) {
	p := &Point{}
	p.SetX(1)
	p.SetY(1)
	if p.X() != 1 {
		t.Fatal()
	}
	if p.Y() != 1 {
		t.Fatal()
	}

}

type TestTypeAssertValer interface {
	X() int
}

type TestTypeAssertVal struct {
	data int
}

func (tt *TestTypeAssertVal) X() int {
	return tt.data
}

func BenchmarkTypeAssert(b *testing.B) {
	var val interface{} = &TestTypeAssertVal{1}
	var sum int

	tmp := val.(TestTypeAssertValer)
	for i := 0; i < b.N; i++ {
		sum += tmp.X()
	}
}

func BenchmarkTypeAssertNormal(b *testing.B) {
	var val = &TestTypeAssertVal{1}
	var sum int

	for i := 0; i < b.N; i++ {
		sum += val.data
	}
}
