package lru

type Node struct {
	pre, next *Node
	val       int
}

type LRU struct {
	m       map[int]*Node
	Root    *Node
	maxsize int
}

func New(size int) *LRU {
	return &LRU{
		m: map[int]*Node{},
		Root: &Node{
			pre:  nil,
			next: nil,
			val:  -1,
		},
		maxsize: size,
	}
}

func (l *LRU) Add(val int) {
	node, ok := l.m[val]
	if ok {
		// 插入表头
		node.pre.next = node.next
		node.next = l.Root.next
		node.pre = l.Root
		if l.Root.next != nil {
			l.Root.next.pre = node
		}
		l.Root.next = node
		l.m[val] = node
	} else {
		// 生成新节点
		node = &Node{
			pre:  l.Root,
			next: l.Root.next,
			val:  val,
		}
		if l.Root.next != nil {
			l.Root.next.pre = node
		}
		l.Root.next = node
		l.m[val] = node
	}
	// 判断容量是否超出限制
	if len(l.m) > l.maxsize {
		cur := l.Root
		for cur.next != nil {
			cur = cur.next
		}
		cur.pre = nil
	}
}
