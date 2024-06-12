package mypackege

type Point struct {
	X, Y int
}

func (p Point)Distance(q Point) int  {
	println(p.X)
	println(p.Y)
	return q.X+q.Y
}