type item struct {
	key int
	value int
}

type MyHashMap struct {
	size int
	buckets [][]*item
}

func Constructor() MyHashMap {
    size := 100
	buckets := make([][]*item, size)

	for i, _ := range buckets {
		buckets[i] = []*item{}
	}

	return MyHashMap{
		size: size,
		buckets: buckets,
	}
}

func (t *MyHashMap) bucketID(key int) int {
	return key%t.size
}

func (t *MyHashMap) Put(key int, value int) {
    ele, ok := t.get(key)
	if ok {
		ele.value = value
		return
	}
	bucketID := t.bucketID(key)
	t.buckets[bucketID] = append(
		t.buckets[bucketID],
		&item{key: key, value: value},
	)
}

func (t *MyHashMap) get(key int) (*item, bool) {
    bucketID := t.bucketID(key)
	for _, item := range t.buckets[bucketID] {
		return item, true
	}
	return nil, false
}

func (t *MyHashMap) Get(key int) int {
    item, ok := t.get(key)
	if !ok {
		return -1
	}
	return item.value
}

func (t *MyHashMap) Remove(key int) {
    bucketID := t.bucketID(key)
	for i, item := range t.buckets[bucketID] {
		if item.key == key {
			t.buckets[bucketID] = append(
				t.buckets[bucketID][:i],
				t.buckets[bucketID][i+1:]...,
			)
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */