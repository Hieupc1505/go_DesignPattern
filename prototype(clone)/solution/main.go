package solution

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Clone() Point {
	return Point{
		p.X,
		p.Y,
	}
}

type Node struct {
	Value    Point
	Children []Point
}

func (n *Node) Clone() Node {
	children := make([]Point, len(n.Children))

	for i := range children {
		children[i] = n.Children[i].Clone()
	}

	return Node{
		Value:    n.Value.Clone(),
		Children: children,
	}
}

func main() {
	p1 := Point{1, 2}
	p2 := p1.Clone()

	p1.X = 3
	fmt.Println(p1, p2)

	n1 := Node{
		Value: p1,
		Children: []Point{
			{4, 5},
			{6, 7},
		},
	}

	n2 := n1.Clone()
	n1.Children[0].X = 9
	fmt.Println(n1, n2) //now n2 is a 'deep copy' from n1

}

// Another example with a 'value object' (ex: Color)
type Color struct {
	R, G, B byte
}

// New color copied from recived color
func (c *Color) ColorWithRed(r byte) Color {
	return Color{r, c.G, c.B}
}
