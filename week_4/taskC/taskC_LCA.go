package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Person struct {
	Parent string
	Depth  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nstr, _ := reader.ReadString('\n')
	n := 0
	fmt.Sscanf(nstr, "%d", &n)

	tree := make(map[string]Person)

	for i := 0; i < n-1; i++ {
		input, _ := reader.ReadString('\n')
		relation := strings.Fields(input)
		child, parent := relation[0], relation[1]
		tree[child] = Person{Parent: parent, Depth: 0}
		if _, exist := tree[parent]; !exist {
			tree[parent] = Person{Parent: "", Depth: 0}
		}
	}

	var root string
	for name, person := range tree {
		if person.Parent == "" {
			root = name
			break
		}
	}

	calculateDepth(tree, root, 0)

	names := make([]string, 0, len(tree))
	for name := range tree {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s %d\n", name, tree[name].Depth)
	}

}

func calculateDepth(tree map[string]Person, name string, depth int) {
	tree[name] = Person{Parent: tree[name].Parent, Depth: depth}
	for child, person := range tree {
		if person.Parent == name {
			calculateDepth(tree, child, depth+1)
		}
	}
}
