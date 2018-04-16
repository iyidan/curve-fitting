package leastsquares

import (
	"errors"

	"github.com/iyidan/curve-fitting/base"
)

// LS Least squares object
type LS struct {
	b1 float64
	b0 float64
}

// New allocate and init a LS with the given points
func New(points []base.Pointer) (base.CurveFittinger, error) {
	if len(points) < 2 {
		return nil, errors.New("point num must gte 2")
	}
	ls := &LS{}
	ls.init(points)
	return ls, nil
}

// GetY implement base.CurveFittinger
func (ls *LS) GetY(x float64) (y float64) {
	return ls.b1*x + ls.b0
}

func (ls *LS) init(points []base.Pointer) {
	var n, sumX, sumY, sumX2, sumXY float64

	n = float64(len(points))

	for i := 0; i < len(points); i++ {
		sumX += points[i].X()
		sumY += points[i].Y()
		sumX2 += points[i].X() * points[i].X()
		sumXY += points[i].X() * points[i].Y()
	}

	println(n, sumX, sumY, sumX2, sumXY)

	ls.b1 = (sumXY*n - sumX*sumY) / (sumX2*n - sumX*sumX)
	ls.b0 = sumY/n - ls.b1*sumX/n
}
