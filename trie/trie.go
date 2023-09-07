package trie

type Node struct {
	IsEnd    bool
	Children map[rune]*Node
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{
		root: &Node{
			IsEnd:    false,
			Children: make(map[rune]*Node),
		}}
}

func (t *Trie) Add(words string) {
	node := t.root
	for _, ch := range words {
		child, ok := node.Children[ch]
		if !ok {
			child = &Node{IsEnd: false, Children: make(map[rune]*Node)}
			node.Children[ch] = child
		}
		node = child
	}
	node.IsEnd = true
}

func (t *Trie) Exits(words string) bool {
	node := t.root
	for _, ch := range words {
		child, ok := node.Children[ch]
		if !ok {
			return false
		}
		node = child
	}
	return true
}
