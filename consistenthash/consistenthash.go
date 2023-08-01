package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes into uint32
type Hash func(data []byte) uint32

// Map contains all hashed keys
type Map struct {
	hash     Hash
	replicas int
	vNodes   []int
	hashMap  map[int]string
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some nodes to the hash map.
func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ {
			vHash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.vNodes = append(m.vNodes, vHash)
			m.hashMap[vHash] = node
		}
	}
	sort.Ints(m.vNodes)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.vNodes) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.vNodes), func(i int) bool {
		return m.vNodes[i] >= hash
	})

	return m.hashMap[m.vNodes[idx%len(m.vNodes)]]
}
