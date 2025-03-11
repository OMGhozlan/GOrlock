package instruct

import (
	"reflect"
	"testing"
)

func assertValueMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertFloatValue(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("testing perimeter of a rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		assertValueMessage(t, got, want)
	})
	t.Run("testing area of a rectangle", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0
		assertFloatValue(t, got, want)
	})
	t.Run("testing area of a rectangle struct", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		got := rec.Area()
		want := 100.0
		assertValueMessage(t, got, want)
	})
	t.Run("testing area of a circle struct", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		assertValueMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {
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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Testing area calculation for Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Testing area calculation for Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Testing area calculation for Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
