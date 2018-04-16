package leastsquares

import (
	"testing"

	"github.com/iyidan/curve-fitting/base"
)

func TestLS(t *testing.T) {
	points := []base.Pointer{&base.Point{Dx: 1, Dy: 1}}
	ls, err := New(points)
	if err == nil || ls != nil {
		t.Fatal()
	}

	points = append(points, &base.Point{Dx: 2, Dy: 2})
	ls, err = New(points)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ls: %#v", ls)

	if ls.GetY(0) != 0 {
		t.Fatal()
	}

	t.Log(3, ls.GetY(3))
	t.Log(4, ls.GetY(4))
	t.Log(5, ls.GetY(5))
}

func TestLSSlice(t *testing.T) {
	points := []float64{0}
	ls, err := NewSlice(points)
	if err == nil || ls != nil {
		t.Fatal()
	}

	points = append(points, 1, 2)
	ls, err = NewSlice(points)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ls: %#v", ls)

	if ls.GetY(-1) != -1 {
		t.Fatal()
	}

	t.Log(3, ls.GetY(3))
	t.Log(4, ls.GetY(4))
	t.Log(5, ls.GetY(5))
}

func TestLSSlice2(t *testing.T) {
	points := [][2]float64{[2]float64{1, 1}}
	ls, err := NewSlice2(points)
	if err == nil || ls != nil {
		t.Fatal()
	}

	points = append(points, [2]float64{2, 2})
	ls, err = NewSlice2(points)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ls: %#v", ls)

	if ls.GetY(0) != 0 {
		t.Fatal()
	}

	t.Log(3, ls.GetY(3))
	t.Log(4, ls.GetY(4))
	t.Log(5, ls.GetY(5))
}
