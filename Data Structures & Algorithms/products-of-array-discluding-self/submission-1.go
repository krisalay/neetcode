func productExceptSelf(nums []int) []int {
	prefixProduct, postfixProduct := 1, 1
	out := make([]int, len(nums))

	// Left to right pass to set prefix product in out
	for i, num := range nums {
		out[i] = prefixProduct
		prefixProduct *= num
	}

	// Right to left pass to multiply postfix product and prefix product
	for i := len(nums)-1; i >= 0; i-- {
		out[i] *= postfixProduct
		postfixProduct *= nums[i]
	}

	return out
}
