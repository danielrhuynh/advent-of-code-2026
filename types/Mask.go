package types

import "math/bits"
type Mask struct {
	Words []uint64
}

func NewMask(n int) Mask {
	words := (n+63)/ 64
	return Mask{Words: make([]uint64, words)}
}

func (m *Mask) Toggle(i int) {
	w := i / 64
	b := uint(i%64)
	m.Words[w] ^= 1 << b
}

func (m *Mask) XOR(other Mask) {
	if len(m.Words) != len(other.Words) {
		panic("main: mismatched word sizes")
	}
	for i := range m.Words {
		m.Words[i] ^= other.Words[i]
	}
}

func (m *Mask) equals(other Mask) bool{
	if len(m.Words) != len(other.Words) {
		return false
	}
	for i := range m.Words {
		if m.Words[i] != other.Words[i] {
			return false
		}
	}
	return true
}

func (m Mask) hammingWeight() int {
	total := 0
	for _, w := range m.Words {
		total += bits.OnesCount64(w)
	}
	return total
}
