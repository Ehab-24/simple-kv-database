package btree

import "bytes"

const (
	BNODE_NODE = 0
	BNODE_LEAF = 1
)

const HEADER = 4
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

type BTree struct {
	root uint64

	get func(uint64) BNode
	new func(BNode) uint64
	del func(uint64)
}

func nodeLookupLE(node BNode, key []byte) uint16 {
	nKeys := node.nKeys()
	found := uint16(0)

	for i := uint16(1); i < nKeys; i++ {
		if cmp := bytes.Compare(node.getKey(i), key); cmp <= 0 {
			found = i
		} else {
			break
		}
	}
	return found
}
