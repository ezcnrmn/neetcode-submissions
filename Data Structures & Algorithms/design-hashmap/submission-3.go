type Node struct {
	Key int
	Val int
	Del bool
}

type MyHashMap struct {
	Arr [10000]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) hash(key int) int {
	return key%10000
}

func (this *MyHashMap) Put(key int, value int) {
	hash := this.hash(key)
	i := hash
	for {
		if this.Arr[i] != nil {
			if this.Arr[i].Key == key || this.Arr[i].Del {
				this.Arr[i].Val = value
				return
			}
		} else {
			this.Arr[i] = &Node{Key: key, Val: value}
			return
		}
		i++
		if i == 10000 {
			i = 0
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	hash := this.hash(key)
	i := hash
	for {
		if this.Arr[i] != nil {
			if this.Arr[i].Key == key && !this.Arr[i].Del {
				return this.Arr[i].Val
			}
		} else {
			break
		}
		i++
		if i == 10000 {
			i = 0
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := this.hash(key)
	i := hash
	for {
		if this.Arr[i] != nil {
			if this.Arr[i].Key == key {
				this.Arr[i].Del = true
			}
		} else {
			break
		}
		i++
		if i == 10000 {
			i = 0
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