//go:build ignore

package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hash_s2 := make(map[rune]int)
	for _, letter := range s2 {
		hash_s2[letter]++
	}

	hash_s1 := make(map[rune]int)
	for _, letter := range s1 {
		hash_s1[letter]++
	}

	for letter, counter := range hash_s2 {
		if targetCounter, ok := hash_s2[letter]; !ok || targetCounter != counter {
			return false
		}
	}
	return true
}

func main() {
	s1 := "racecar"
	s2 := "carrace"
	fmt.Println(isAnagram(s1, s2))

	s3 := "jar"
	s4 := "jam"
	fmt.Println(isAnagram(s3, s4))
}
