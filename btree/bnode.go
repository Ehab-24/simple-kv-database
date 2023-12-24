package btree

import "encoding/binary"

type BNode struct {
	data []byte
}

func init() {
	nodeimax := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	assert(nodeimax <= BTREE_PAGE_SIZE)
}

func offsetPos(node BNode, idx uint16) uint16 {
	assert(1 <= idx && idx <= node.nKeys())
	return HEADER + 8*node.nKeys() + 2*(idx-1)
}

func kvPos(n BNode, idx uint16) uint16 {
	assert(idx <= n.nKeys())
	return HEADER + 8*n.nKeys() + 2*n.nKeys() + n.getOffset(idx)
}

func (n BNode) bType() uint16 {
	return binary.LittleEndian.Uint16(n.data[0:2])
}

func (n BNode) nKeys() uint16 {
	return binary.LittleEndian.Uint16(n.data[2:4])
}

func (n BNode) setHeader(bType uint16, nKeys uint16) {
	binary.LittleEndian.PutUint16(n.data[0:2], bType)
	binary.LittleEndian.PutUint16(n.data[2:4], nKeys)
}

func (n BNode) getPtr(idx uint16) uint64 {
	assert(idx < n.nKeys())
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint64(n.data[pos:])
}

func (n BNode) setPtr(idx uint16, val uint64) {
	assert(idx < n.nKeys())
	pos := HEADER + 8*idx
	binary.LittleEndian.PutUint64(n.data[pos:], val)
}

func (n BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(n.data[offsetPos(n, idx):])
}

func (n BNode) setOffset(idx uint16, val uint16) {
	binary.LittleEndian.PutUint16(n.data[offsetPos(n, idx):], val)
}

func (n BNode) getKey(idx uint16) []byte {
	assert(idx < n.nKeys())
	pos := kvPos(n, idx)
	klen := binary.LittleEndian.Uint16(n.data[pos:])
	return n.data[pos+4:][:klen]
}

func (n BNode) getVal(idx uint16) []byte {
	assert(idx < n.nKeys())
	pos := kvPos(n, idx)
	klen := binary.LittleEndian.Uint16(n.data[pos:])
	vlen := binary.LittleEndian.Uint16(n.data[pos+2:])
	return n.data[pos+4+klen:][:vlen]
}

func (n BNode) size() uint16 {
	return kvPos(n, n.nKeys())
}
