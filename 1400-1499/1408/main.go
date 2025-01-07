package main

// KMP algo (https://ru.wikipedia.org/wiki/Алгоритм_Кнута_—_Морриса_—_Пратта)

func stringMatching(words []string) []string {
	var out []string
	for i := range words {
		for j := range words {
			if j == i {
				continue
			}
			if len(words[i]) > len(words[j]) {
				continue
			}
			if contains(words[j], words[i]) >= 0 {
				out = append(out, words[i])
				break
			}
		}
	}
	return out
}

func contains(t, s string) int {
	if len(t) == len(s) {
		for i := range t {
			if t[i] == s[i] {
				continue
			} else {
				return -1
			}
		}
		return 0
	}
	if len(t) < len(s) {
		return -1
	}
	if s == "" {
		return 0
	}
	idx := -1
	f := findPrefix(s)
	k := 0
	for i := range t {
		if k > 0 && s[k] != t[i] {
			k = f[k-1]
		}
		if s[k] == t[i] {
			k++
		}
		if k == len(s) {
			idx = i - len(s) + 1
			break
		}
	}
	return idx
}

func findPrefix(in string) []int {
	k := 0
	out := make([]int, 0)
	out = append(out, k)
	pos := 1
	for {
		if pos >= len(in) {
			break
		}
		if in[pos] == in[k] {
			k++
			out = append(out, k)
			pos++
			continue
		}
		if k == 0 {
			pos++
			out = append(out, k)
			continue
		}
		k = findPrefixInt(in[0:k])
	}
	return out
}

func findPrefixInt(in string) int {
	k := 0
	pos := 1
	for {
		if pos >= len(in) {
			break
		}
		if in[pos] == in[k] {
			k++
			pos++
			continue
		}
		if k == 0 {
			pos++
			continue
		}
		k = findPrefixInt(in[0:k])
	}
	return k
}
