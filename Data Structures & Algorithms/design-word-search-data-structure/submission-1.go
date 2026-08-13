type Node struct {
	Children [26]*Node
	IsFinal  bool
}

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{IsFinal: true}}
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.Root
	for _, char := range word {
		if cur.Children[char-'a'] == nil {
			cur.Children[char-'a'] = &Node{}
		}
		cur = cur.Children[char-'a']
	}
	cur.IsFinal = true
}

type StackItem struct {
	Node  *Node
	Index int
}

func (this *WordDictionary) Search(word string) bool {
	stack := []StackItem{{Node: this.Root, Index: 0}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Index == len(word) {
			return cur.Node.IsFinal
		}

		char := word[cur.Index]
		if char == '.' {
			for _, child := range cur.Node.Children {
				if child != nil {
					stack = append(stack, StackItem{Node: child, Index: cur.Index+1})
				}
			}
			continue
		} 

		if cur.Node.Children[char-'a'] == nil {
			continue
		}
		stack = append(stack, StackItem{Node: cur.Node.Children[char-'a'], Index: cur.Index+1})
	}
	return false
}
