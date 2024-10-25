package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertEquals(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertEquals(t, got, tt.want)
	}
}

func assertEquals(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
