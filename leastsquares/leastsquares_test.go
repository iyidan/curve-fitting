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
