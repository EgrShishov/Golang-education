package main

import "fmt"

// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
func main() {
	var s string
	fmt.Scan(&s)

	_map := make(map[rune]int)

	for _, char := range s {
		_map[char]++
	}
	var ans string
	for char, value := range _map {
		if value == 1 {
			ans += string(char)
		}
	}

	fmt.Print(ans)
}
