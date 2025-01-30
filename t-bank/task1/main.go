package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Буферизованный ввод и вывод
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	fmt.Fscan(in, &str) // Считываем строку

	// Находим индексы R и M
	var idxR, idxM = -1, -1
	for i, char := range str {
		if char == 'R' {
			idxR = i
		} else if char == 'M' {
			idxM = i
		}
	}

	// Проверяем условие
	if idxR < idxM {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	fmt.Fscan(in, &str)

	for _, char := range str {
		if char == 'R' {
			fmt.Fprint(out, "Yes")
			return
		} else if char == 'M' {
			fmt.Fprint(out, "No")
			return
		}
	}
} */
