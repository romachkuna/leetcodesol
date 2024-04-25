package main

func longestSubarray(nums []int) int {
	leftPointer, rightPointer := 0, 0

	deleted := 0

	for ; rightPointer < len(nums); rightPointer++ {
		deleted += 1 - nums[rightPointer]
		if deleted > 1 {
			deleted -= 1 - nums[leftPointer]
			leftPointer++
		}
	}
	return rightPointer - leftPointer - 1

}

func longestAltitude(gain []int) int {
	currentMax := 0
	step := 0

	for i := 0; i < len(gain); i++ {
		step = step + gain[i]

		if step > currentMax {
			currentMax = step
		}
	}

	return currentMax
}

func pivotIndex(nums []int) int {
	totalSum := 0
	leftSum := 0

	for _, e := range nums {
		totalSum += e
	}

	for i, e := range nums {
		if leftSum == totalSum-leftSum-e {
			return i
		}
		leftSum += e
	}

	return -1
}
func findDifference(nums1 []int, nums2 []int) [][]int {

	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)

	result := make([][]int, 2)

	for _, e := range nums1 {
		nums1Map[e] = true
	}

	for _, e := range nums2 {
		nums2Map[e] = true
	}

	for key, _ := range nums1Map {
		_, ok := nums2Map[key]
		if !ok {
			result[0] = append(result[0], key)
		}
	}

	for key, _ := range nums2Map {
		_, ok := nums1Map[key]
		if !ok {
			result[1] = append(result[1], key)
		}
	}

	return result
}

func uniqueOccurrences(arr []int) bool {
	occuranceMap := make(map[int]int)

	for _, e := range arr {
		_, ok := occuranceMap[e]

		if ok {
			occuranceMap[e] = occuranceMap[e] + 1
		} else {
			occuranceMap[e] = 1
		}
	}

	frequenceMap := make(map[int]bool)

	for _, value := range occuranceMap {
		_, ok := frequenceMap[value]
		if ok {
			return false
		} else {
			frequenceMap[value] = true
		}
	}

	return true
}
