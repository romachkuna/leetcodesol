package main

func main() {
	longestSubarray([]int{1, 0, 1, 0, 1, 1})
}

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
