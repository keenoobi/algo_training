package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Person struct {
	Parent     string
	Descendant int
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
		tree[child] = Person{Parent: parent, Descendant: 0}
		if _, exist := tree[parent]; !exist {
			tree[parent] = Person{Parent: "", Descendant: 0}
		}
	}

	childrenMap := make(map[string][]string)
	for child, person := range tree {
		parent := person.Parent
		childrenMap[parent] = append(childrenMap[parent], child)
	}

	for name := range tree {
		calculateDescendants(tree, childrenMap, name)
	}

	names := make([]string, 0, len(tree))
	for name := range tree {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s %d\n", name, tree[name].Descendant)
	}

}

func calculateDescendants(tree map[string]Person, childrenMap map[string][]string, name string) int {
	if tree[name].Descendant != 0 {
		return tree[name].Descendant
	}
	descendants := 0
	for _, child := range childrenMap[name] {
		descendants += 1 + calculateDescendants(tree, childrenMap, child)
	}
	tree[name] = Person{Parent: tree[name].Parent, Descendant: descendants}
	return descendants
}
