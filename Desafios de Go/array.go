package main

func TwoSum(nums []int, target int) []int {
	numMap := map[int]int{}

	for i, num := range nums {
		diferenca := target - num
		if index, ok := numMap[diferenca]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return nil
}
