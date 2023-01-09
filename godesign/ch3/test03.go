package main

import "fmt"

type Point struct {
	X, Y int
}
type address struct {
	hostname string
	port     int
}

func main() {
	//p := Point{}
	//fmt.Println(p)
	//fmt.Println(Scale(Point{1, 2}, 5))
	//pp := &Point{1, 2}
	pp := new(Point)
	*pp = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(*pp)

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)
	a := map[address]int{
		address{port: 111}: 1,
	}
	fmt.Println(a)

}
func Scale(p Point, factor int) Point {
	return Point{X: p.X * factor, Y: p.Y * factor}
}
