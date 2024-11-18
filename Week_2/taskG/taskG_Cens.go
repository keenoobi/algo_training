package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fstr := strings.Fields(input)
	N, _ := strconv.Atoi(fstr[0])
	C, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')

	runes := []rune(scndstr)

	left, right, cntA, cntB := 0, 0, 0, 0
	maxLen := 0
	roughness := 0

	for ; right < N; right++ {
		if runes[right] == 'a' {
			cntA++
		} else if runes[right] == 'b' {
			cntB++
			roughness += cntA
		}

		for ; roughness > C; left++ {
			if runes[left] == 'a' {
				cntA--
				roughness -= cntB
			} else if runes[left] == 'b' {
				cntB--
			}
		}
		maxLen = max(maxLen, right-left+1)
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, maxLen)
	writer.Flush()
}
