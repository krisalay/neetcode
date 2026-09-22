func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	freqBucket := make([][]int, len(nums)+1)

	for _, num := range nums {
		freqMap[num]++
	}

	for num, freq := range freqMap {
		freqBucket[freq] = append(freqBucket[freq], num)
	}

	var res []int
	for i := len(freqBucket)-1; i > 0; i-- {
		for _, num := range freqBucket[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}