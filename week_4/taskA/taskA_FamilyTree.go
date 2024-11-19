package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Person struct {
	Child  string
	Parent string
	Height int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nstr, _ := reader.ReadString('\n')
	var n int
	fmt.Sscanf(nstr, "%d", &n)

	relations := make([][]string, n-1)
	for i := 0; i < n-1; i++ {
		input, _ := reader.ReadString('\n')
		relation := strings.Fields(input)
		relations[i] = relation
	}

	tree := buildTree(n-1, relations)

	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Child < tree[j].Child
	})

	for _, person := range tree {
		fmt.Printf("%s %d\n", person.Child, person.Height)
	}
}

func buildTree(n int, relations [][]string) []Person {
	people := make([]Person, n)
	childrens := make(map[string][]string)

	for i, relation := range relations {
		child, parent := relation[0], relation[1]
		childrens[relation[0]] = append(childrens[relation[0]], relation[1])
		people[i] = Person{Child: child, Parent: parent, Height: -1}
	}

	var root string
	for _, people := range relations {
		if _, exist := childrens[people[1]]; exist == false {
			root = people[1]
			break
		}
	}

	people = append(people, Person{Child: root, Parent: "", Height: -1})

	var calculateHeight func(name string, height int)
	calculateHeight = func(name string, height int) {
		for i := range people {
			if people[i].Child == name {
				people[i].Height = height
				break
			}
		}
		for i := range people {
			if people[i].Parent == name {
				calculateHeight(people[i].Child, height+1)
			}
		}
	}

	calculateHeight(root, 0)

	return people
}
