package tempconv

type MathFunction interface {
	Calculate(arg float64) float64
	GetName() string
}

type CToF struct {
	Degree float64
	Name   string
}

func (c CToF) Calculate(degree float64) float64 {
	return degree*1.8 + 32
}

func (c CToF) GetName() string {
	return c.Name
}

type FToC struct {
	Degree float64
	Name   string
}

func (f FToC) Calculate(degree float64) float64 {
	return 0.56 * (degree - 32)
}

func (f FToC) GetName() string {
	return f.Name
}

type CToK struct {
	Degree float64
	Name   string
}

func (c CToK) Calculate(degree float64) float64 {
	return degree + 273.15
}

func (c CToK) GetName() string {
	return c.Name
}

type KToC struct {
	Degree float64
	Name   string
}

func (c KToC) Calculate(degree float64) float64 {
	return degree - 273.15
}

func (c KToC) GetName() string {
	return c.Name
}

type FToK struct {
	Degree float64
	Name   string
}

func (c FToK) Calculate(degree float64) float64 {
	return (degree-32)*5/9 + 273.15
}

func (c FToK) GetName() string {
	return c.Name
}

type KToF struct {
	Degree float64
	Name   string
}

func (c KToF) Calculate(degree float64) float64 {
	return (degree-273.15)*9/5 + 32
}

func (c KToF) GetName() string {
	return c.Name
}
