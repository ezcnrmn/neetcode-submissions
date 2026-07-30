type MyHashSet struct {
	buckets [10000]*node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (s *MyHashSet) Add(key int) {
	cur := s.buckets[key%10000]
	var prev *node
	for cur != nil {
		if cur.val == key {
			return
		}
		prev = cur
		cur = cur.next
	}
	if prev != nil {
		prev.next = &node{val: key}
	} else {
		s.buckets[key%10000] = &node{val: key}
	}
}

func (s *MyHashSet) Remove(key int) {
	cur := s.buckets[key%10000]
	var prev *node
	for cur != nil {
		if cur.val == key {
			break
		}
		prev = cur
		cur = cur.next
	}
	if cur != nil {
		return
	}
	if prev != nil {
		prev.next = cur.next
	} else {
		s.buckets[key%10000] = nil
	}
}

func (s *MyHashSet) Contains(key int) bool {
	cur := s.buckets[key%10000]
	for cur != nil {
		if cur.val == key {
			return true
		}
		cur = cur.next
	}
	return false
}

// My node for list

type node struct {
	val  int
	next *node
}