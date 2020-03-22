package main

import "fmt"

func solve(str string) int {
	var max = -1
	var current = 0

	for _, chr := range str {
		if chr == 'a' || chr == 'i' || chr == 'e' || chr == 'o' || chr == 'u' {
			if current > max {
				max = current
			}

			current = 0
		} else {
			current += int(chr - 'a') + 1
		}
	}

	if current > max {
		max = current
	}

	return max
}

func main() {
	fmt.Println(solve("zodiacs"))
	fmt.Println(solve("strength"))
}
