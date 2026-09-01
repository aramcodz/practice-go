package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{5.0, 10.0}

	result := rect.Perimiter()
	expect := 30.0

	if expect != result {
		t.Errorf("expected perimeter %.2f, actual perim %.2f", expect, result)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("checkArea helper - expected Area %g, result Area: %g", expected, result)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		expected := 55.0
		rectangle := Rectangle{5, 11}
		checkArea(t, rectangle, expected)
	})

	t.Run("circles"), func(t *testing.T)  {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, t, expected)	
	})
}

//Previous Version without using "Shape" interface
// t.Run("rectangles", func(t *testing.T) {
// 	rect := Rectangle{5.0, 11.0}
// 	result := rect.Area()

// 	expect := 55.0

// 	if expect != result {
// 		t.Errorf("expected rectangle Area %.2f, actual Area %.2f", expect, result)
// 	}
// })

// t.Run("circles", func(t *testing.T) {
// 	circle := Circle{10}
// 	actual := circle.Area()

// 	expected := 314.1592653589793
// 	if expected != actual {
// 		t.Errorf("expected Area %.2f, actual Area %.2f", expected, actual)
// 	}
// })
