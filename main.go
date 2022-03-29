package main

func main() {
	nums := []int{1, 3, 4, 2, 2}
	findDuplicate(nums)
}

func findDuplicate(nums []int) int {
	var result = 0
	var mapValues = make(map[int]int)

	for index := range nums {
		value := nums[index]

		if _, ok := mapValues[value]; ok {
			result = value
		} else {
			mapValues[nums[index]] = 1
		}
	}

	return result
}
