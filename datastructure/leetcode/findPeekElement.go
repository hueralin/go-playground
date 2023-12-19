package leetcode

// FindPeekElement 162 寻找峰值
func FindPeekElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		//if nums[mid] > nums[mid+1] {
		//	right = mid
		//} else {
		//	left = mid + 1
		//}
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
