package main
import "fmt"
import "math"

type Node11 struct {
	item string
	left *Node11
	right *Node11
}

func main() {

	a := Node11 {
		item: "A",
		left: nil,
		right: nil,
	}

  c := Node11 {
		item: "C",
		left: nil,
		right: nil,
	}

  d := Node11 {
		item: "D",
		left: &c,
		right: nil,
	}

	b := Node11 {
		item: "B",
		left: &a,
		right: &d,
	}

	g := Node11 {
		item: "G",
		left: nil, 
		right: nil,
	} 
	f := Node11 {
		item: "F",
		left: nil, 
		right: &g,
	}

	e := Node11 {
		item: "E",
		left: &b,
		right: &f,
	}

  fmt.Println(e)
  depth(&e, 1)

}
var depthG = 0.0
func depth(root *Node11, d int) {
  
  if root == nil {
    return 
  }

  if root.left == nil && root.right == nil {
    fmt.Println("I am here", root.item)
      depthG = math.Max(float64(depthG), float64(d))
      fmt.Println(math.Max(float64(depthG), float64(d)))

  }

  depth(root.left, d + 1)
  depth(root.right, d + 1)
}