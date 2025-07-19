package bptree

import "sort"

// BPlusTreeIndex - Simplified B+Tree (simulated with sorted slice for demo)
// Note: This is in-memory and not scalable beyond moderate sizes.
// For scaling to millions+ of records, a disk-backed or distributed B+Tree/LSM-tree is required.
// Benefits: Range queries are efficient, multiple attributes can be indexed.
// Drawbacks: More complex I/O and caching, higher write amplification (LSM) vs. memory costs (B+Tree).
// Tradeoffs: LSM-trees give faster writes but slower range queries; B+Trees provide faster range reads.
type BPlusTreeIndex struct {
	keys   []int
	values map[int][]string // maps size or timestamp to list of file IDs
}

// NewBPlusTreeIndex creates a new index.
func NewBPlusTreeIndex() *BPlusTreeIndex {
	return &BPlusTreeIndex{
		keys:   []int{},
		values: make(map[int][]string),
	}
}

// Insert adds a key-value pair to the index.
func (b *BPlusTreeIndex) Insert(key int, fileID string) {
	if _, ok := b.values[key]; !ok {
		b.keys = append(b.keys, key)
		sort.Ints(b.keys)
	}
	b.values[key] = append(b.values[key], fileID)
}

// RangeLookup returns all file IDs for keys within [min, max].
func (b *BPlusTreeIndex) RangeLookup(min, max int) []string {
	var result []string
	for _, k := range b.keys {
		if k >= min && k <= max {
			result = append(result, b.values[k]...)
		}
	}
	return result
}
