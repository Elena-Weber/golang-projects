package main

import "fmt"

type shape interface { // all getArea functions that return float64 belong to this type
	getArea() float64
}

type triangle struct{ // this is like a hash in Ruby
	base float64
	height float64
}

type square struct{ // this is like a hash in Ruby
	side float64
}

func main(){
	// these are like local vars of types defined above
	t := triangle{base: 10, height: 10}
	s := square{side: 10}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}