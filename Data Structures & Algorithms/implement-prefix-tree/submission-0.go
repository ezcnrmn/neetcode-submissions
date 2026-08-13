type Node struct {
	Children [26]*Node
	Char     rune
	IsFinal  bool
}

type PrefixTree struct {
	Head *Node
}

func Constructor() PrefixTree {
	return PrefixTree{Head: &Node{IsFinal: true}}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.Head
	for _, char := range word {
		if cur.Children[char-'a'] == nil {
			cur.Children[char-'a'] = &Node{Char: char}
		}
		cur = cur.Children[char-'a']
	}
	cur.IsFinal = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.Head
	for _, char := range word {
		if cur.Children[char-'a'] == nil {
			return false
		}
		cur = cur.Children[char-'a']
	}
	return cur.IsFinal
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.Head
	for _, char := range prefix {
		if cur.Children[char-'a'] == nil {
			return false
		}
		cur = cur.Children[char-'a']
	}
	return true
}
