package main

import "fmt"

// является ли строка палиндромом
func main() {
	var s string
	fmt.Scan(&s)
	ans := true

	str := []rune(s)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			ans = false
			break
		}
	}

	if ans == true {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}

}
