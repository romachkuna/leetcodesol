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
