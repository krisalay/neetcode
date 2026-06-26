func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	countBucket := make([][]int, len(nums)+1)
	for num, count := range freqMap {
		countBucket[count] = append(countBucket[count], num)
	}

	result := make([]int, k)
	for i := len(countBucket)-1; i >= 0; i-- {
		if k == 0 {
			return result
		}
		for _, num := range countBucket[i] {
			result[k-1] = num
			k--
		}
	}
	return result
}
