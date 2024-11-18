package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Algorithm struct {
	Index      int
	Interest   int
	Usefulness int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numAlgorithmsStr, _ := reader.ReadString('\n')
	numAlgorithmsStr = strings.TrimSpace(numAlgorithmsStr)
	numAlgorithms, _ := strconv.Atoi(numAlgorithmsStr)

	interestsStr, _ := reader.ReadString('\n')
	interestsStr = strings.TrimSpace(interestsStr)
	interests := strings.Fields(interestsStr)
	interestValues := make([]int, numAlgorithms)
	for i := 0; i < numAlgorithms; i++ {
		interestValues[i], _ = strconv.Atoi(interests[i])
	}

	usefulnessStr, _ := reader.ReadString('\n')
	usefulnessStr = strings.TrimSpace(usefulnessStr)
	usefulness := strings.Fields(usefulnessStr)
	usefulnessValues := make([]int, numAlgorithms)
	for i := 0; i < numAlgorithms; i++ {
		usefulnessValues[i], _ = strconv.Atoi(usefulness[i])
	}

	moodStr, _ := reader.ReadString('\n')
	moodStr = strings.TrimSpace(moodStr)
	mood := strings.Fields(moodStr)
	moodValues := make([]int, numAlgorithms)
	for i := 0; i < numAlgorithms; i++ {
		moodValues[i], _ = strconv.Atoi(mood[i])
	}

	algorithms := make([]Algorithm, numAlgorithms)
	for i := 0; i < numAlgorithms; i++ {
		algorithms[i] = Algorithm{Index: i + 1, Interest: interestValues[i], Usefulness: usefulnessValues[i]}
	}

	sort.Slice(algorithms, func(i, j int) bool {
		if algorithms[i].Interest == algorithms[j].Interest {
			return algorithms[i].Usefulness > algorithms[j].Usefulness
		}
		return algorithms[i].Interest > algorithms[j].Interest
	})

	algorithmsByUsefulness := make([]Algorithm, numAlgorithms)
	copy(algorithmsByUsefulness, algorithms)

	sort.Slice(algorithmsByUsefulness, func(i, j int) bool {
		if algorithmsByUsefulness[i].Usefulness == algorithmsByUsefulness[j].Usefulness {
			return algorithmsByUsefulness[i].Interest > algorithmsByUsefulness[j].Interest
		}
		return algorithmsByUsefulness[i].Usefulness > algorithmsByUsefulness[j].Usefulness
	})

	selected := make([]bool, numAlgorithms)
	result := make([]int, numAlgorithms)

	for i := 0; i < numAlgorithms; i++ {
		var chosen Algorithm
		if moodValues[i] == 1 {

			for j := 0; j < numAlgorithms; j++ {
				if !selected[algorithmsByUsefulness[j].Index-1] {
					chosen = algorithmsByUsefulness[j]
					selected[chosen.Index-1] = true
					break
				}
			}
		} else {

			for j := 0; j < numAlgorithms; j++ {
				if !selected[algorithms[j].Index-1] {
					chosen = algorithms[j]
					selected[chosen.Index-1] = true
					break
				}
			}
		}
		result[i] = chosen.Index
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, index := range result {
		fmt.Fprintf(writer, "%d ", index)
	}
	writer.Flush()
}
