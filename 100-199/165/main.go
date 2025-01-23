package main

import (
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1Items := strings.Split(version1, ".")
	v2Items := strings.Split(version2, ".")

	maxItems := len(v1Items)
	if len(v2Items) > maxItems {
		maxItems = len(v2Items)
	}

	for i := 0; i < maxItems; i++ {
		if len(v1Items) <= i && len(v2Items) <= i {
			return 0
		}
		v1Current := "0"
		v2Current := "0"
		if len(v1Items) > i {
			v1Current = v1Items[i]
		}
		if len(v2Items) > i {
			v2Current = v2Items[i]
		}
		result := compareVersionParts(v1Current, v2Current)
		if result != 0 {
			return result
		}
	}
	return 0
}

func compareVersionParts(a, b string) int {
	var aPos = 0
	var bPos = 0

	for {
		if aPos >= len(a) {
			break
		}
		if a[aPos] == '0' {
			aPos++
		} else {
			break
		}
	}

	for {
		if bPos >= len(b) {
			break
		}
		if b[bPos] == '0' {
			bPos++
		} else {
			break
		}
	}

	if ((len(b) - 1) - bPos) > ((len(a) - 1) - aPos) {
		return -1
	} else if ((len(a) - 1) - aPos) > ((len(b) - 1) - bPos) {
		return 1
	}

	if bPos == len(b) || aPos == len(a) {
		return 0
	}

	for {
		if aPos == len(a) || bPos == len(b) {
			return 0
		}
		if a[aPos] > b[bPos] {
			return 1
		} else if b[bPos] > a[aPos] {
			return -1
		}
		aPos++
		bPos++
	}
}
