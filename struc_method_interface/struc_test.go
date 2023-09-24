package struc

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Permiter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestTableArea(t *testing.T) {

	//here we are defining a slice of structs then we fill it in with the cases
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{

		{name: "rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				//%#v prints our structs with values in the field
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
