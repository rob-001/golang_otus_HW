package main

import "fmt"

func main() {
	var n int
	var result string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	fmt.Println(result)
}
