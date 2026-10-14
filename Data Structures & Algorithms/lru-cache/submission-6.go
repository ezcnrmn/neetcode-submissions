type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Map  map[int]*Node
	Head *Node
	Tail *Node
	Cap  int
	Len  int
}

func (this *LRUCache) cut(node *Node) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) putToHead(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head

	this.Head.Next.Prev = node
	this.Head.Next = node
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail

	return LRUCache{
		Map: make(map[int]*Node,capacity),
		Head: head,
		Tail: tail,
		Cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		this.cut(node)
		this.putToHead(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Map[key]; ok {
		node.Val = value
		this.cut(node)
		this.putToHead(node)
		return
	}

	node := &Node{Key: key, Val: value}
	this.Map[key] = node
	
	this.putToHead(node)
	this.Len++
	if this.Cap < this.Len {
		lastNode := this.Tail.Prev
		this.cut(lastNode)
		delete(this.Map, lastNode.Key)
		this.Len--
	}
}
