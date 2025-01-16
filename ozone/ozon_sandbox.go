package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isValidFormat(s string) bool {

	if len(s) > 0 && (s[0] == ' ' || s[len(s)-1] == ' ') {
		return false
	}

	if strings.Contains(s, "  ") {
		return false
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		input1 := scanner.Text()

		scanner.Scan()
		input2 := scanner.Text()

		if !isValidFormat(input2) {
			fmt.Fprintln(out, "no")
			continue
		}

		str1 := strings.Split(input1, " ")
		str2 := strings.Split(input2, " ")

		if len(str1) != n || len(str2) != n {
			fmt.Fprintln(out, "no")
			continue
		}

		array := make([]int, n)
		var err error

		for j := 0; j < n; j++ {
			array[j], err = strconv.Atoi(str1[j])
			if err != nil {
				fmt.Fprintln(out, "no")
				break
			}
		}
		if err != nil {
			continue
		}

		sort.Ints(array)

		equal := true
		for k := 0; k < n; k++ {
			val, err := strconv.Atoi(str2[k])
			if err != nil || array[k] != val {
				fmt.Fprintln(out, "no")
				equal = false
				break
			}
		}

		if equal {
			fmt.Fprintln(out, "yes")
		}
	}
}
