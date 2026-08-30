package main

type Rectange struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectange) float64 {
	return (rect.Width + rect.Height) * 2
}

func Area(rect Rectange) float64 {
	return rect.Height * rect.Width
}
