func productExceptSelf(nums []int) []int {
	prefixArray := make([]int, len(nums))
	postfixArray := make([]int, len(nums))
	out := make([]int, len(nums))
	prefixProduct, postfixProduct := 1, 1

	// Build prefix array
	for i, num := range nums {
		prefixArray[i] = prefixProduct
		prefixProduct *= num
	}
	// Build postfix array
	for i := len(nums)-1; i >= 0; i-- {
		postfixArray[i] = postfixProduct
		postfixProduct *= nums[i]
	}
	// Build output
	for i, num := range prefixArray {
		out[i] = num * postfixArray[i]
	}
	return out
}
