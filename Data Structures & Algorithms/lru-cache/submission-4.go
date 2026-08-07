type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Cap  int
	Head *Node
	Tail *Node
	Map  map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Head: head,
		Tail: tail,
		Cap:  capacity,
		Map:  make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) addToHead(node *Node) {
	this.Head.Next.Prev = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.addToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		node.Val = value
		this.remove(node)
		this.addToHead(node)
		return
	}

	if len(this.Map) >= this.Cap {
		lastNode := this.Tail.Prev
		delete(this.Map, lastNode.Key)
		this.remove(lastNode)
	}

	node = &Node{Key: key, Val: value}
	this.addToHead(node)
	this.Map[key] = node
}
