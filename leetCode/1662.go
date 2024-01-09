package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 string
	for _, v := range word1 {
		s1 += v
	}
	for _, v := range word2 {
		s2 += v
	}
	return s1 == s2
}


