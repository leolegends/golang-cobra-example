package app

// Calc struct
type Calc struct {
	A float64
	B float64
}

//Sum return float.
func (c Calc) Sum() float64 {
	return c.A + c.B
}

func NewCalc() Calc {
	return Calc{}
}
