package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1

	for _, ch := range input {
		// fmt.Println("index:", i, "Run:", ch)

		// min, max := 'A', 'Z'

		// if ch >= min && ch <= max {
		// 	answer++
		// }

		// if ch < 69 && ch > 37 {
		// 	answer++
		// }

		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}

	}

	fmt.Println(answer)
}
