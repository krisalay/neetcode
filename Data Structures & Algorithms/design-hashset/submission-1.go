type MyHashSet struct {
	buckets [][]int
	size int
}

func Constructor() MyHashSet {
    size := 100
	buckets := make([][]int, size)

	for i, _ := range buckets {
		buckets[i] = []int{}
	}

	return MyHashSet{
		size: size,
		buckets: buckets,
	}
}

func (this *MyHashSet) getBucketIdx(key int) int {
	return key%this.size
}

func (this *MyHashSet) Add(key int) {
    if this.Contains(key) {
		return
	}

	bucketIdx := this.getBucketIdx(key)
	this.buckets[bucketIdx] = append(this.buckets[bucketIdx], key)
}

func (this *MyHashSet) Remove(key int) {
    bucketIdx := this.getBucketIdx(key)
	for i, d := range this.buckets[bucketIdx] {
		if d == key {
			this.buckets[bucketIdx] = append(
				this.buckets[bucketIdx][:i],
				this.buckets[bucketIdx][i+1:]...,
			)
			break
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    bucketId := this.getBucketIdx(key)

	for _, d := range this.buckets[bucketId] {
		if d == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 