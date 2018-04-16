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

	p := make([][2]float64, len(points))
	for _, v := range points {
		p = append(p, [2]float64{v.X(), v.Y()})
	}

	ls := &LS{}
	ls.init(p)
	return ls, nil
}

// NewSlice allocate and init a LS with the given slice points
func NewSlice(points []float64) (base.CurveFittinger, error) {
	if len(points) < 2 {
		return nil, errors.New("point num must gte 2")
	}

	ls := &LS{}
	ls.initWithSlice(points)
	return ls, nil
}

// NewSlice2 same as NewSlice
func NewSlice2(points [][2]float64) (base.CurveFittinger, error) {
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

func (ls *LS) init(points [][2]float64) {
	var n, sumX, sumY, sumX2, sumXY float64
	n = float64(len(points))
	for i := 0; i < len(points); i++ {
		sumX += points[i][0]
		sumY += points[i][1]
		sumX2 += points[i][0] * points[i][0]
		sumXY += points[i][0] * points[i][1]
	}
	ls.genParams(n, sumX, sumY, sumX2, sumXY)
}

func (ls *LS) initWithSlice(points []float64) {
	var n, sumX, sumY, sumX2, sumXY float64
	n = float64(len(points))
	for i, y := range points {
		x := float64(i)
		sumX += x
		sumY += y
		sumX2 += x * x
		sumXY += x * y
	}
	ls.genParams(n, sumX, sumY, sumX2, sumXY)
}

func (ls *LS) genParams(n, sumX, sumY, sumX2, sumXY float64) {
	ls.b1 = (sumXY*n - sumX*sumY) / (sumX2*n - sumX*sumX)
	ls.b0 = sumY/n - ls.b1*sumX/n
}
