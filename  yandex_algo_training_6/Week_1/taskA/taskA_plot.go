package main

import "fmt"

func task1() {
	var coord [6]int
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d\n", &coord[i])
	}

	if coord[5] > coord[3] { // north
		fmt.Print("N")
	} else if coord[5] < coord[1] { // south
		fmt.Print("S")
	}

	if coord[4] > coord[2] { // east
		fmt.Print("E")
	} else if coord[4] < coord[0] { // west
		fmt.Print("W")
	}
}
