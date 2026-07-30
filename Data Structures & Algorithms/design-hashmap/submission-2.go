type MyHashMap struct {
	buckets [10000]*node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (m *MyHashMap) Put(key int, value int) {
	cur := m.buckets[key%10000]
	var prev *node
	for cur != nil {
		if cur.key == key {
			cur.val = value
			return
		}
		prev = cur
		cur = cur.next
	}
	if prev != nil {
		prev.next = &node{key: key, val: value}
	} else {
		m.buckets[key%10000] = &node{key: key, val: value}
	}
}

func (m *MyHashMap) Get(key int) int {
	cur := m.buckets[key%10000]
	for cur != nil {
		if cur.key == key {
			return cur.val
		}
		cur = cur.next
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	cur := m.buckets[key%10000]
	var prev *node
	for cur != nil {
		if cur.key == key {
			break
		}
		prev = cur
		cur = cur.next
	}
	if prev != nil {
		prev.next = cur.next
	} else if cur != nil {
		m.buckets[key%10000] = cur.next
	} else {
		m.buckets[key%10000] = nil
	}
}

type node struct {
	key  int
	val  int
	next *node
}