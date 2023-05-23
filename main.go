package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}

func main() {
	_map := make(map[int]int)
	for i := 0; i < 10; i++ {
		var c = 0
		fmt.Scan(&c)
		if value, inMap := _map[c]; inMap {
			fmt.Println(value)
		} else { //i love dima
			_map[c] = work(c)
			fmt.Println(c)
		}
	}

}
