package struc

// a struct is a named collection of fields where you can store data
type Rectangle struct {
	Width  float64
	Height float64
}

func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return (rectangle.Width * rectangle.Height)
}
