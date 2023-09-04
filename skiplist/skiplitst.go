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

/**
时间复杂度，O(lg n),每一层是常数次比较，  空间复杂度 O(n)
主要理清楚查找，插入思路，什么时候搜索终止，什么时候找到插入的节点，单链表在插入和删除时都比较依赖于先找到前一个节点，
所以针对于跳表来讲，每一层都有一个最小的节点会比较容易做删除和搜索相关的工作。

*/

func Constructor() Skiplist {
	return Skiplist{head: &Node{val: -1, down: nil, next: nil}}
}

// 搜索过程，当前节点下一节点 > target ， 从当前节点下一节点开始寻找，
// 当前节点下一节点 <  target ，向当前节点右节点开始寻找
// 当前节点下一个节点等于target 直接返回
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

// 插入过程中如何找到节点插入的位置，找到次小于target的位置
// 插入过程，找到节点，插入到终止节点右侧,还需要存储每一个层的节点

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

// 删除过程， 从头节点开始遍历每一层，
// 当前指针的下一个指针== target 删除 ,循环删除

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
	return deleted
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
