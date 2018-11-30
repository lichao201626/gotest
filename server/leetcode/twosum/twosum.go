package twosum

import ()

// TwoSum function find two index added 48ms
/* func TwoSum(nums []int, target int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
} */

// define a struct to store the index and value 24ms
/* // Bucket...
type Bucket struct {
	Index int
	Value int
}

// TwoSum function find two index added
func TwoSum(nums []int, target int) []int {
	sum := 0
	var temp []Bucket
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(temp); j++ {
			if nums[i] == temp[j].Value {
				return []int{i, temp[j].Index}
			}
		}
		sum = nums[i] + nums[i+1]
		if sum == target {
			return []int{i, i + 1}
		} else {
			v := target - nums[i]
			bk := Bucket{Index: i, Value: v}
			temp = append(temp, bk)
		}
	}
	return []int{}
} */

// using go map 8ms
func TwoSum(nums []int, target int) []int {
	result := make([]int, 2)
	num_index_map := make(map[int]int)
	for i, num := range nums {
		if idx, existed := num_index_map[target-num]; existed {
			result[0] = idx
			result[1] = i
			return result
		}
		num_index_map[num] = i
	}
	return result
}
