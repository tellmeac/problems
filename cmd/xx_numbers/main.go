package main

func main() {
	missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
}

func missingNumber(nums []int) int {
	nums = append(nums, -1)
	n := len(nums)

	for idx := 0; idx < n; idx++ {
		if nums[idx] != idx {
			otherIdx := nums[idx]
			if otherIdx == 1 {
				continue
			}

			nums[idx], nums[otherIdx] = nums[otherIdx], nums[idx]
		}
	}

	for idx := range nums {
		if nums[idx] != idx {
			return idx
		}
	}

	return n
}
