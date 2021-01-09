package longest

func FindLongest(words []string) []string {
	var longest []string
	for i, word := range words {
		longest = append(longest, word)
		for j, word2 := range words {
			if i != j && Compare(word, word2) && !Contains(longest, word2) {
				longest = append(longest, word2)
			}
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
