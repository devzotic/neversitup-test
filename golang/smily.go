package main

import "regexp"

func CountSmilyFace(faces []string) int {
	if len(faces) == 0 {
		return 0
	}

	facePattern := `[:;][-~]?[)D]`

	regexp := regexp.MustCompile(facePattern)

	count := 0
	for _, face := range faces {
		if regexp.MatchString(face) {
			count++
		}
	}

	return count
}
