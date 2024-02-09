package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		checkArea(t, rectangle, 72.0)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//
//}

func TestArea(t *testing.T) {
	// Table-Driven Tests Approach

	// Create a slice of structs, where each element of the slice must:
	// Have a shape property
	// Have a want property
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// Add all the shapes here
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
