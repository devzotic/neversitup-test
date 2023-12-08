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
		arr[start], arr[i] = arr[i], arr[start] // => 0, 0 => 0, 1 => 0, 2 // abc (start) == abc =>  bac (abc) = > cba
		Manipulate(string(arr), start+1, out)   // bac (start) => 1, 1 => 1, 2 //
		arr[start], arr[i] = arr[i], arr[start]
	}
}
