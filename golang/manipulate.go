package main

import "slices"

func Manipulate(s string, start int, out *[]string) {
	if start == len(s)-1 {
		if !slices.Contains(*out, s) {
			*out = append(*out, string(s))
			return
		}
	}

	arr := []rune(s)
	for i := start; i < len(s); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		Manipulate(string(arr), start+1, out)
		arr[start], arr[i] = arr[i], arr[start]
	}
}
