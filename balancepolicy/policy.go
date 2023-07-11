package balancepolicy

import "sync"

type Policy interface {
	AddNode(addr string, nodeName string)
	PickNode(key string) string
}

type RoundRobin struct {
	lock     sync.Mutex
	index    int64
	nodeInfo map[string]string
	nodes    []string
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{
		index:    0,
		nodeInfo: map[string]string{},
		nodes:    make([]string, 0),
	}
}

func (r *RoundRobin) AddNode(addr string, nodeName string) {
	r.nodeInfo[nodeName] = addr
	r.nodes = append(r.nodes, nodeName)
}

func (r *RoundRobin) PickNode(key string) string {
	r.lock.Lock()
	defer r.lock.Unlock()
	num := r.index % int64(len(r.nodes))
	r.index++
	if r.index < 0 {
		// 防止溢出
		r.index = 0
	}
	return r.nodeInfo[r.nodes[num]]
}
