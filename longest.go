package longest

import "fmt"

func FindLongest(words []string) []string {
	var longest []string
	for _, word := range words {
		newLongest := []string{word}
		for _, word2 := range words {
			if Compare(word, word2) && !Contains(newLongest, word2) {
				newLongest = append(newLongest, word2)
				word = word2
			}
		}
		fmt.Println(newLongest)
		if len(newLongest) > len(longest) {
			longest = newLongest
		}
	}

	return longest
}

func Compare(first, last string) bool {
	fl := string(first[len(first)-1])
	ll := string(last[0])
	return fl == ll
}

func Contains(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}
