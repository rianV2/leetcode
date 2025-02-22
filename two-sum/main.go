package main

import "fmt"

func twoSum(nums []int, target int) []int {
	bucket := make(map[int]int)

	for i, v := range nums {
		find := target - v
		if ie, exist := bucket[find]; exist {
			return []int{ie, i}
		}

		bucket[v] = i
	}

	return []int{}
}

func main() {
	// input := [][]int{
	// 	[]int{2},
	// }
	fmt.Println(twoSum([]int{3, 3, 1}, 4))
}
