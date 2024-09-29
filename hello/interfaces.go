package main

type Shape interface {
	area() float32
}

type Triangle struct {
	height float32
	base   float32
}

type Square struct {
	length float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}
func (s Square) area() float32 {
	return s.length * s.length
}

func (r Rectangle) area() float32 {
	return r.breadth * r.length
}
