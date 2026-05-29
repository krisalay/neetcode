func sortColors(nums []int) {
    var zeros int
	var ones int
	var twos int

	for _, num := range nums {
		if num == 0 {
			zeros++
		}
		if num == 1 {
			ones++
		}
		if num == 2 {
			twos++
		}
	}

	write := 0
	for range zeros {
		nums[write] = 0
		write++
	}
	for range ones {
		nums[write] = 1
		write++
	}
	for range twos {
		nums[write] = 2
		write++
	}
}
