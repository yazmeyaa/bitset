package bitset

type BitSet struct {
	bits []uint32
}

func NewBitSet(elements int) *BitSet {
	numBits := (elements + 31) / 32
	return &BitSet{
		bits: make([]uint32, numBits),
	}
}

func (bs *BitSet) Set(index int) {
	bs.bits[index/32] |= 1 << (index % 32)
}

func (bs *BitSet) Clear(index int) {
	bs.bits[index/32] &^= 1 << (index % 32)
}

func (bs *BitSet) IsSet(index int) bool {
	return bs.bits[index/32]&(1<<(index%32)) != 0

}
