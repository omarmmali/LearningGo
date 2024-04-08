package shapes

import "testing"

func TestRectangles(t *testing.T) {
	got := Rectangle{10.0, 10.0}.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestCircles(t *testing.T) {
	got := Circle{10}.Perimeter()
	want := 62.83185307179586

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		checkArea(t, Rectangle{12.0, 6.0}, 72.0)
	})

	t.Run("Circle", func(t *testing.T) {
		checkArea(t, Circle{10}, 314.1592653589793)
	})
}

func checkArea(t *testing.T, shape Shape, want float64) {
	got := shape.Area()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
