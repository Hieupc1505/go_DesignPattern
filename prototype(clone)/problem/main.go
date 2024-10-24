package problem

import "fmt"

type Point struct {
	X, Y int
}

type Node struct {
	Value    Point
	Children []Point
}

func main() {
	p1 := Point{1, 2}
	p2 := p1

	p1.X = 3
	fmt.Println(p1.X, p2.X)

	n1 := Node{
		Value: p1,
		Children: []Point{
			{4, 5},
			{6, 7},
		},
	}

	n2 := n1
	n1.Children[0].X = 9
	fmt.Println(n1, n2)
}
