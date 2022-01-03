/* using map to find the first non repeating character but map does not maintain order, so result may vary */

package main

import (
	"fmt"
)

func main() {
	s := "priypdharsshini"
	m := make(map[string]int)

	for _, key := range s {
		m[string(key)]++
	}

	for key, val := range m {
		if val == 1 {
			fmt.Println("The first non repeated characted is")
			fmt.Println(key)
			break
		}

	}
}
