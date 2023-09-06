package heap

type Heap struct {
	arr []int
}

func HeapInsert(arr []int) *Heap {
	h := &Heap{arr: make([]int, 0, len(arr))}
	for _, num := range arr {
		h.Insert(num)
	}
	return h
}

func Heapify(arr []int) *Heap {
	h := &Heap{arr: arr}
	lastNotLeaf := len(arr)/2 - 1
	for i := lastNotLeaf; i >= 0; i-- {
		h.ShiftDown(i)
	}
	return h
}

func (h *Heap) Insert(num int) {
	h.arr = append(h.arr, num)
	h.ShiftUp(len(h.arr) - 1)
}

// 从标号为index的节点开始做shifUp操作
func (h *Heap) ShiftUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if h.arr[parent] < h.arr[index] {
		swap(h.arr, parent, index)
		h.ShiftUp(parent)
	}
}

// 删除并返回根节点
func (h *Heap) Pop() int {
	num := h.arr[0]
	swap(h.arr, 0, len(h.arr)-1)
	h.arr = h.arr[:len(h.arr)-1]
	h.ShiftDown(0)
	return num
}

// 从标号为index的节点开始做shifDown操作
func (h *Heap) ShiftDown(index int) {
	left := index*2 + 1
	right := index*2 + 2
	if left < len(h.arr) && right < len(h.arr) {
		if h.arr[left] >= h.arr[right] && h.arr[left] > h.arr[index] {
			swap(h.arr, left, index)
			h.ShiftDown(left)
		}
		if h.arr[right] > h.arr[left] && h.arr[right] > h.arr[index] {
			swap(h.arr, right, index)
			h.ShiftDown(right)
		}
	}
	if left >= len(h.arr) {
		return
	}
	if right >= len(h.arr) {
		if h.arr[left] > h.arr[index] {
			swap(h.arr, left, index)
			h.ShiftDown(left)
		}
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
