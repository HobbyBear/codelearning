package skiplist

import "math/rand"

type Node struct {
	val  int
	next *Node
	down *Node
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {
	return Skiplist{head: &Node{val: -1, down: nil, next: nil}}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for cur != nil {
		next := cur.next
		down := cur.down
		if cur.val == target {
			return true
		}
		if next != nil && next.val == target {
			return true
		}
		if next != nil && next.val > target {
			cur = down
			continue
		}
		if next != nil && next.val < target {
			cur = next
		}
		if next == nil {
			cur = down
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	arr := make([]*Node, 0)
	cur := this.head
	for cur != nil {
		next := cur.next
		down := cur.down
		if next != nil && next.val >= num {
			arr = append(arr, cur)
			cur = down
			continue
		}
		if next != nil && next.val < num {
			cur = next
		}
		if next == nil {
			arr = append(arr, cur)
			cur = down
		}
	}
	inserted := true
	var down *Node
	for i := len(arr) - 1; i >= 0; i-- {
		if inserted {
			next := arr[i].next
			arr[i].next = &Node{val: num, next: next, down: down}
			inserted = rand.Intn(1000)%2 == 0
			down = arr[i].next
		}
	}
	// 添加新层
	if inserted {
		this.head = &Node{val: -1, down: this.head, next: &Node{val: num, next: nil, down: down}}
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur := this.head
	deleted := false
	for cur != nil {
		next := cur.next
		down := cur.down
		if next != nil && next.val == num {
			cur.next = next.next
			next.down = nil
			next.next = nil
			cur = down
			deleted = true
		}
		if next != nil && next.val > num {
			cur = down
			continue
		}
		if next != nil && next.val < num {
			cur = next
		}
		if next == nil {
			cur = down
		}
	}
	if deleted {
		// 清理掉删除节点后没有元素的空层
		cur = this.head
		for cur.next == nil {
			this.head = cur.down
			cur = this.head
		}
	}
	return deleted
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
