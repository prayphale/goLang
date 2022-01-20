package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	// switch {
	// case strings.IndexRune(alphabetLower, ch) >= 0:
	// 	ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
	// case strings.IndexRune(alphabetUpper, ch) >= 0:
	// 	ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
	// default:
	// 	ret = ret + string(ch)
	// }

	// if strings.IndexRune(alphabetStr, ch) >= 0 {
	// 	ret = ret + string(rotate(ch, delta, alphabet))
	// } else {
	// 	ret = ret + string(rotate(ch, delta, alphabet))
	// }
	// }
	// a := rotateData('Z', 2, 3)
	fmt.Println(string(ret))

}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotateData(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotateData(r, 'a', delta)
	}
	return r
}

func rotateData(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]

	// idx := -1
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }

	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx >= len(key) {
	// 		idx = 0
	// 	}
	// }

	// return key[idx]
}
