package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	seen := make(map[rune]bool)

	for _, r := range strings.ToLower(s) {
		if seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}
