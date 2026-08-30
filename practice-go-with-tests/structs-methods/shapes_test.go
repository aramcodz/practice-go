package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectange{5.0, 10.0}
	result := Perimeter(rect)
	expect := 30.0

	if expect != result {
		t.Errorf("expected perimeter %.2f, actual perim %.2f", expect, result)
	}
}

func TestArea(t *testing.T) {
	rect := Rectange{5.0, 11.0}
	result := Area(rect)

	expect := 55.0

	if expect != result {
		t.Errorf("expected Area %.2f, actual Area %.2f", expect, result)
	}
}
