package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	origColor := image[sr][sc]
	image[sr][sc] = color

	if isValid(image, sr-1, sc) {
		if image[sr-1][sc] == origColor {
			floodFill(image, sr-1, sc, color)
		}
	}
	if isValid(image, sr, sc+1) {
		if image[sr][sc+1] == origColor {
			floodFill(image, sr, sc+1, color)
		}
	}
	if isValid(image, sr+1, sc) {
		if image[sr+1][sc] == origColor {
			floodFill(image, sr+1, sc, color)
		}
	}
	if isValid(image, sr, sc-1) {
		if image[sr][sc-1] == origColor {
			floodFill(image, sr, sc-1, color)
		}
	}
	return image
}

func isValid(image [][]int, sr, sc int) bool {
	if sr < 0 || sr > len(image)-1 || sc < 0 || sc > len(image[0])-1 {
		return false
	}
	return true
}
