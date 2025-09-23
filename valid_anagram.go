package main
import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash_t := make(map[rune]int)
	for _, letter := range t {
		hash_t[letter]++
	}

	hash_s := make(map[rune]int)
	for _, letter := range s {
		hash_s[letter]++
	}

	for letter, counter := range hash_t {
		if targetCounter, ok := hash_s[letter]; !ok || targetCounter != counter {
			return false
		}
	}
	return true
}

func main() {
	s1 := "racecar"
	t1 := "carrace"
	fmt.Println(isAnagram(s1, t1))

	s2 := "jar"
	t2 := "jam"
	fmt.Println(isAnagram(s2, t2))
}
