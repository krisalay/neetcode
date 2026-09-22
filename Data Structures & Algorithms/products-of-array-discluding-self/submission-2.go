func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefixProduct := getPrefixProduct(nums)
	product := 1
	for i := len(nums)-1; i >= 0; i-- {
		res[i] = prefixProduct[i] * product
		product *= nums[i]
	}
	return res
}

func getPrefixProduct(nums []int) []int {
	res := make([]int, len(nums))
	product := 1
	for i, num := range nums {
		res[i] = product
		product *= num
	}
	return res
}
