package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sqr Square) End() Point {
	a := int(sqr.a)
	x := sqr.start.x + a
	y := sqr.start.y + a

	return Point{x: x, y: y}
}

func (sqr Square) Area() uint {
	return sqr.a * sqr.a
}

func (sqr Square) Perimeter() uint {
	return sqr.a * 4
}
