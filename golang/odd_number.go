package main

func findOdd(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for k, v := range countMap {
		if v%2 != 0 {
			return k, true
		}
	}

	return 0, false
}
