package longest

//Will need to create multiple possible combinations to piece together
//The most efficient way to keep track of this is to use maps
type matchingFirstLetter map[string][]string
type matchingLastLetter map[string][]string

/*
FindLongest takes a slice of strings and returns a subset of them
(potentially the same length of the initial slice)
with the last letter of the first string matched to the first letter of the next string
*/

var fl = make(matchingFirstLetter)
var ll = make(matchingLastLetter)

func FindLongest(words []string) []string {
	for _, word := range words {
		fl[string(word[0])] = append(fl[string(word[0])], word)
		ll[string(word[len(word)-1])] = append(ll[string(word[len(word)-1])], word)
	}
	var longest []string
	//range over words
	for _, word := range words {
		longest = CreateNewLongest(word, words, longest)
	}
	return longest
}

func CreateNewLongest(word string, words, longest []string) []string {
	newLongest := []string{word}
	last := len(newLongest) - 1
	for i := 0; i < len(words); i++ {
		l := string(newLongest[last][last])
		arr, ok := fl[l]
		if ok {
			//at this point, it may make sense to make a recursive function, but I'm not sure how yet.
			for _, word := range arr {
				newLongest = append(newLongest, word)
				//fix recursive solution to not do a stack overflow
				return CreateNewLongest(string(newLongest[last][last]), words, longest)
			}
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
