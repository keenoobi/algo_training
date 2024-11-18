package main

import (
	"fmt"
)

// getPriorityMap создает карту приоритета для лексикографического порядка
func getPriorityMap(w string) map[rune]int {
	priority := make(map[rune]int)
	for i, ch := range w {
		priority[ch] = i + 1
	}
	return priority
}

// buildMinSequence строит минимальную лексикографическую ПСП
func buildMinSequence(n int, w, s string) string {
	priority := getPriorityMap(w)
	result := []rune(s)
	stack := []rune{}

	sumStack := 0

	// Инициализируем стек на основе строки `s`
	for _, ch := range s {
		if ch == '(' || ch == '[' {
			stack = append(stack, ch) // Добавляем открывающие скобки в стек
			sumStack += priority[ch]
		}
	}

	for len(result) < n {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top == '(' {
				result = append(result, ')')
				sumStack -= priority[')']
			} else {
				result = append(result, ']')
				sumStack -= priority[']']
			}
			stack = stack[:len(stack)-1]
		}
	}

	// // Добавляем скобки для завершения до длины n
	// for len(result) < n {
	// 	// Отладочный вывод для каждого шага
	// 	fmt.Println("Stack:", string(stack), "Result:", string(result))

	// 	// Проверяем, можем ли еще добавлять открывающие скобки, чтобы сохранить баланс
	// 	if len(stack) > 0 {
	// 		// Если стек не пустой, закрываем последнюю открытую скобку
	// 		top := stack[len(stack)-1]
	// 		if top == '(' {
	// 			result = append(result, ')')
	// 		} else {
	// 			result = append(result, ']')
	// 		}
	// 		stack = stack[:len(stack)-1]
	// 	} else {
	// 		if priority['('] < priority['['] {
	// 			result = append(result, '(')
	// 			stack = append(stack, '(')
	// 		} else {
	// 			result = append(result, '[')
	// 			stack = append(stack, '[')
	// 		}
	// 	}
	// }

	return string(result)
}

func main() {
	// Входные данные
	var n int
	var w, s string
	fmt.Scan(&n, &w, &s)

	// Строим минимальную ПСП
	fmt.Println(buildMinSequence(n, w, s))
}
