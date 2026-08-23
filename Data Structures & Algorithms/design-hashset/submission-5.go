type Node struct {
    Val  int
    Next *Node
}

type MyHashSet struct {
    Map [10000]*Node
}

func Constructor() MyHashSet {
    var dummies [10000]*Node
    for i := range dummies {
        dummies[i] = &Node{}
    }
    return MyHashSet{
        Map: dummies,
    }
}

func (this *MyHashSet) Add(key int) {
    cur := this.Map[key%10000]
    for cur.Next != nil {
        if cur.Next.Val == key {
            return
        }
        cur = cur.Next
    }
    cur.Next = &Node{Val: key}
}

func (this *MyHashSet) Remove(key int) {
    cur := this.Map[key%10000]
    for cur.Next != nil {
        if cur.Next.Val == key {
            cur.Next = cur.Next.Next
            return
        }
        cur = cur.Next
    }
}

func (this *MyHashSet) Contains(key int) bool {
    cur := this.Map[key%10000]
    for cur.Next != nil {
        if cur.Next.Val == key {
            return true
        }
        cur = cur.Next
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
 