package main

import (
	"fmt"
)

func main() {
	s := "priypdharsshini"
	m := make(map[string]int)
	for _, val := range s {
		m[string(val)]++
	}

	for _, val := range s {
		if m[string(val)] == 1 {
			fmt.Println("The first non repeating character is")
			fmt.Println(string(val))
			break
		}
	}
}
