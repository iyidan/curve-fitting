package base

// Pointer (x,y)
type Pointer interface {
	X() float64
	Y() float64
	SetX(x float64)
	SetY(y float64)
}

// CurveFittinger is a interface which implemented by the undering algorithms
type CurveFittinger interface {
	// GetY get a curve-fitted value by x
	GetY(x float64) (y float64)
}

// Point implement Pointer
type Point struct {
	Dx float64
	Dy float64
}

// X implement Pointer
func (p *Point) X() float64 {
	return p.Dx
}

// Y implement Pointer
func (p *Point) Y() float64 {
	return p.Dy
}

// SetX implement Pointer
func (p *Point) SetX(x float64) {
	p.Dx = x
}

// SetY implement Pointer
func (p *Point) SetY(y float64) {
	p.Dy = y
}
