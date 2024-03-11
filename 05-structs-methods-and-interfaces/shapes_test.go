package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	got := Area(rectangle)
	// 	want := 72.0

	// 	// %.2f 2 decimal points
	// 	// %g long float
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := Area(circle)
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTableDriven(t *testing.T) {
	// we define a slice of anonymous structs with fields shape and want
	// and then fill the slice with cases
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	// we can then loop over this slice and run the tests
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

// you can optionally name the fields
// now our tests - rather, the list of test cases - make assertions of truth about shapes and their areas
// instead of having to manually look through the cases to find out which case failed
// we can change our error message into %#v
// this will print out our struct with the values in its field, so we can see the properties being tested
func TestAreaTableDrivenImproved(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the t.Run test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

// Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

// Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

// Adding methods so you can add functionality to your data types and so you can implement interfaces

// Table driven tests to make your assertions clearer and your test suites easier to extend & maintain

// We are now starting to define our own types. In statically typed languages like Go, being able to design your own types is essential for building software that is easy to understand, to piece together and to test.

// Interfaces are a great tool for hiding complexity away from other parts of the system.
