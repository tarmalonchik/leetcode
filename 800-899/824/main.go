package main

var vowelMp = map[byte]interface{}{
	'a': nil,
	'e': nil,
	'i': nil,
	'o': nil,
	'u': nil,
}

func toGoatLatin(sentence string) string {
	sentenceByte := []byte(sentence)
	pos1 := 0
	pos2 := 0
	counter := 1
	out := make([]byte, 0)

	for {
		pos1, pos2 = nextWord(sentenceByte, pos1)
		if pos1 == -1 {
			break
		}
		out = append(out, convertWord(sentenceByte[pos1:pos2], counter)...)
		out = append(out, []byte(" ")...)
		counter++
		pos1 = pos2
	}
	out = out[:len(out)-1]
	return string(out)
}

func nextWord(in []byte, pos int) (pos1, pos2 int) {
	for {
		if pos > len(in)-1 {
			return -1, -1
		}
		if in[pos] == ' ' {
			pos++
			continue
		}
		break
	}
	pos1 = pos
	pos2 = pos
	for {
		if pos2 > len(in)-1 {
			return pos1, pos2
		}
		if in[pos2] == ' ' {
			break
		}
		pos2++
	}
	return pos1, pos2
}

func convertWord(word []byte, counter int) []byte {
	newWord := make([]byte, len(word), len(word)+3+counter)
	copy(newWord, word)
	if isVowel(word[0]) {
		newWord = append(newWord, []byte("ma")...)
	} else {
		newWord = append(newWord, newWord[0])
		newWord = append(newWord, []byte("ma")...)
		newWord = newWord[1:]
	}
	for i := 0; i < counter; i++ {
		newWord = append(newWord, 'a')
	}
	return newWord
}

func isVowel(in byte) bool {
	if _, ok := vowelMp[toLower(in)]; ok {
		return true
	}
	return false
}

func toLower(in byte) byte {
	if in >= 65 && in <= 90 {
		return in + 32
	}
	return in
}
