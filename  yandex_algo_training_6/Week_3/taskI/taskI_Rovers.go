package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Rover struct {
	direction int
	time      int
	index     int
}

func main() {
	var N, a, b int
	fmt.Scan(&N)
	fmt.Scan(&a, &b)

	rovers := make([]Rover, N)
	for i := 0; i < N; i++ {
		var d, t int = 0, 0
		fmt.Scan(&d, &t)
		rovers[i] = Rover{direction: d, time: t, index: i}
	}

	// Сортируем роверы по времени приезда
	sort.Slice(rovers, func(i, j int) bool {
		return rovers[i].time < rovers[j].time
	})

	// Определяем главную и второстепенную дороги
	mainRoads := map[int]bool{a: true, b: true}

	// Очередь для роверов, готовых проехать перекресток
	queue := list.New()
	current_time := 0
	result := make([]int, N)

	for _, rover := range rovers {
		// Обновляем текущее время
		current_time = max(current_time, rover.time)

		// Добавляем ровер в очередь
		queue.PushBack(rover)

		// Обрабатываем очередь
		for queue.Len() > 0 {
			// Находим роверы, которые могут проехать перекресток
			readyToPass := []Rover{}
			for e := queue.Front(); e != nil; e = e.Next() {
				r := e.Value.(Rover)
				if r.time <= current_time {
					readyToPass = append(readyToPass, r)
				}
			}

			if len(readyToPass) == 0 {
				break
			}

			// Сортируем по приоритету (сначала главная дорога)
			sort.Slice(readyToPass, func(i, j int) bool {
				ri, rj := readyToPass[i], readyToPass[j]
				if mainRoads[ri.direction] != mainRoads[rj.direction] {
					return mainRoads[ri.direction]
				}
				return ri.direction < rj.direction
			})

			// Пропускаем роверы
			for _, r := range readyToPass {
				result[r.index] = current_time
				for e := queue.Front(); e != nil; e = e.Next() {
					if e.Value.(Rover) == r {
						queue.Remove(e)
						break
					}
				}
			}

			// Увеличиваем время на 1, так как роверы проехали перекресток
			current_time++
		}
	}

	// Выводим результаты
	for _, time := range result {
		fmt.Println(time)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
