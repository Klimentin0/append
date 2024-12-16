package main

type cost struct {
	day   int
	value float64
}

// Проверяем число дней, по максимальному номеру дня
func findMaxDay(costs []cost) int {
	maxDay := costs[0].day
	for i := 0; i < len(costs); i++ {
		if maxDay < costs[i].day {
			maxDay = costs[i].day
		}
	}
	return maxDay
}

func getCostsByDay(costs []cost) []float64 {
	// Используем функцию для поиска самого последнего дня
	maxDay := findMaxDay(costs)

	// Создаём слайс с длиной в кол-стве отображаемых дней
	dailyCost := make([]float64, maxDay)
	// bubble сортировка
	var sortedSlice []cost
	n := len(costs)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if costs[j].day > costs[j+1].day {
				costs[j], costs[j+1] = costs[j+1], costs[j]
			}
		}
	}
	// Делаем список по одному одному дню со сложенными vlaues
	for i := 0; i < len(costs); i++ {
		found := false
		for j := 0; j < len(sortedSlice); j++ {
			if costs[i].day == sortedSlice[j].day {
				sortedSlice[j].value += costs[i].value
				found = true
				break
			}
		}
		if !found {
			sortedSlice = append(sortedSlice, cost{day: costs[i].day, value: costs[i].value})
		}
	}

	for i := 0; i < len(sortedSlice); i++ {
		item := sortedSlice[i]

		dailyCost = append(dailyCost, item.value)
	}
	return dailyCost
}
