package struc

import "math"

// In Go interface resolution is implicit.
// If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

// a struct is a named collection of fields where you can store data
type Rectangle struct {
	Width  float64
	Height float64
}

// a method is a function with a reciever, we can only call methods on types
type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// the structure for a method is func (receiverName ReceiverType) MethodName(args)
// It is convention in go for the reciever to be the first letter of the type
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
