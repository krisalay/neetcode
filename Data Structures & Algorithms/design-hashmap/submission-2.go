type node struct {
	key, value int
	next *node
}

type MyHashMap struct {
	store []*node
	bucketSize int
}

func Constructor() MyHashMap {
	bucketSize := 1000
    return MyHashMap{
		bucketSize: bucketSize,
		store: make([]*node, bucketSize),
	}
}

func (this *MyHashMap) bucketID(key int) int {
	return key%this.bucketSize
}

func (this *MyHashMap) Put(key int, value int) {
    bucketID := this.bucketID(key)
	curr := this.store[bucketID]
	newNode := &node{key: key, value: value}
	if curr == nil {
		this.store[bucketID] = newNode
		return
	}
	this.Remove(key)
	newNode.next = this.store[bucketID]
	this.store[bucketID] = newNode
}

func (this *MyHashMap) Get(key int) int {
    bucketID := this.bucketID(key)
	curr := this.store[bucketID]
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    bucketID := this.bucketID(key)
	curr := this.store[bucketID]
	if curr == nil {
		return
	}
	if curr.key == key {
		this.store[bucketID] = curr.next
		return
	}
	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */