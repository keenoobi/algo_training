package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (nd *Node) Insert(value int) *Node {
	if nd == nil {
		fmt.Println("DONE")
		return NewNode(value)
	}
	if value < nd.Value {
		nd.Left = nd.Left.Insert(value)
	} else if value > nd.Value {
		nd.Right = nd.Right.Insert(value)
	} else {
		fmt.Println("ALREADY")
	}
	return nd
}

func (nd *Node) Search(value int) *Node {
	if nd == nil {
		fmt.Println("NO")
		return nd
	} else if nd.Value == value {
		fmt.Println("YES")
		return nd
	}
	if value < nd.Value {
		return nd.Left.Search(value)
	}
	return nd.Right.Search(value)
}

func (nd *Node) PrintTreeWithDepth(depth int) {
	if nd == nil {
		return
	}

	nd.Left.PrintTreeWithDepth(depth + 1)

	for i := 0; i < depth; i++ {
		fmt.Print(".")
	}

	fmt.Println(nd.Value)

	nd.Right.PrintTreeWithDepth(depth + 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var root *Node

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		str := strings.Fields(input)

		switch str[0] {
		case "ADD":
			val, _ := strconv.Atoi(str[1])
			root = root.Insert(val)
		case "SEARCH":
			val, _ := strconv.Atoi(str[1])
			root.Search(val)
		case "PRINTTREE":
			root.PrintTreeWithDepth(0)
		}

	}
}
