package main

import"strings"


func replaceWords(dictionary []string, sentence string) string {
	words := strings.Fields(sentence)

	rootMap := make(map[string]bool)
	for _, root := range dictionary {
		rootMap[root] = true
	}

	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if rootMap[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}


	return strings.Join(words, " ")
}