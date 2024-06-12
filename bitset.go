package bitset

import "sync"

type BitSet struct {
	bits []uint32
	mx   sync.RWMutex
}

func NewBitSet(elements int) *BitSet {
	numBits := (elements + 31) / 32
	return &BitSet{
		bits: make([]uint32, numBits),
	}
}

func (bs *BitSet) Set(index int) {
	bs.mx.Lock()
	defer bs.mx.Unlock()

	bs.bits[index/32] |= 1 << (index % 32)
}

func (bs *BitSet) Clear(index int) {
	bs.mx.Lock()
	defer bs.mx.Unlock()

	bs.bits[index/32] &^= 1 << (index % 32)
}

func (bs *BitSet) IsSet(index int) bool {
	bs.mx.RLock()
	defer bs.mx.RUnlock()

	return bs.bits[index/32]&(1<<(index%32)) != 0

}
