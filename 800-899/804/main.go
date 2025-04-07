package main

func uniqueMorseRepresentations(words []string) int {
	letters := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	mp := make(map[string]interface{})
	for i := range words {
		word := make([]byte, 0)
		for j := range words[i] {
			word = append(word, toMorse(letters, words[i][j])...)
		}
		mp[string(word)] = nil
	}
	return len(mp)
}

func toMorse(mapping []string, char byte) []byte {
	char -= 97
	return []byte(mapping[int(char)])
}
