package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Parent string
	Depth  int
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	reader := bufio.NewReader(file)

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

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		elements := strings.Fields(input)
		if len(elements) != 2 {
			break
		}
		elem1, elem2 := elements[0], elements[1]

		lca := findLCA(tree, elem1, elem2)
		fmt.Println(lca)
	}
}

func findLCA(tree map[string]Person, elem1, elem2 string) string {
	path1 := getPathToRoot(tree, elem1)
	path2 := getPathToRoot(tree, elem2)

	i := 0
	for i < len(path1) && i < len(path2) && path1[i] == path2[i] {
		i++
	}
	return path1[i-1]
}

func getPathToRoot(tree map[string]Person, elem string) []string {
	var path []string
	for elem != "" {
		path = append(path, elem)
		elem = tree[elem].Parent
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
