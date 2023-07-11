package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type ConsistentHash struct {
	nodes      map[uint32]string
	keys       []uint32
	replicates int
}

func New(replicate int) *ConsistentHash {
	return &ConsistentHash{
		nodes:      make(map[uint32]string),
		keys:       make([]uint32, 0),
		replicates: replicate,
	}
}

func (c *ConsistentHash) AddNodes(node string) {
	for i := 0; i <= c.replicates; i++ {
		nodename := fmt.Sprintf("%s#%d", node, i)
		hashKey := crc32.ChecksumIEEE([]byte(nodename))
		c.nodes[hashKey] = nodename
		c.keys = append(c.keys, hashKey)
	}
	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
}

func (c *ConsistentHash) GetNode(key string) string {
	hashKey := crc32.ChecksumIEEE([]byte(key))
	nodekeyIndex := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hashKey
	})
	if nodekeyIndex == len(c.keys) {
		nodekeyIndex = 0
	}
	return c.nodes[c.keys[nodekeyIndex]]
}
