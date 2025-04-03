package main

func toLowerCase(s string) string {
	b := []byte(s)

	for i := range b {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] += 32
		}
	}
	return string(b)
}
