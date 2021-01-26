package longest

import "fmt"

//Will need to create multiple possible combinations to piece together
//The most efficient way to keep track of this is to use maps
//For each word, check which words have matching first letter
//Then see which of those have matching last letter
//Keep coming up with longer and longer iterations doing this back and forth
type matchingFirstLetter map[string][]string
type matchingLastLetter map[string][]string

/*
FindLongest takes a slice of strings and returns a subset of them
(potentially the same length of the initial slice)
with the last letter of the first string matched to the first letter of the next string
*/
func FindLongest(words []string) []string {
	fmt.Println(words)
	var longest []string
	fl := make(matchingFirstLetter)
	ll := make(matchingLastLetter)
	for _, word := range words {
		fl[string(word[0])] = append(fl[string(word[0])], word)
		ll[string(word[len(word)-1])] = append(ll[string(word[len(word)-1])], word)
	}
	// for _, word := range words {
	// 	longest = CreateNewLongest(word, words, longest)
	// }
	fmt.Println(fl)
	fmt.Println(ll)
	return longest
}

func CreateNewLongest(word string, words, longest []string) []string {
	newLongest := []string{word}
	for _, word2 := range words {
		if Compare(word, word2) && !Contains(newLongest, word2) {

			newLongest = append(newLongest, word2)
			word = word2
		}
	}
	if len(newLongest) >= len(longest) {
		return newLongest
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
