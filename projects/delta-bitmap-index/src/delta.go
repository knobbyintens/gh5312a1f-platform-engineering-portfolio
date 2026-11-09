package bitmapindex

import "sort"

// Block is a delta-encoded run of sorted row identifiers. Base is the first
// identifier in the block and Deltas[i] holds the gap to the next identifier.
type Block struct {
	Base   uint32
	Deltas []uint16
}

// EncodeDelta builds delta blocks from a sorted slice of identifiers. The
// first identifier of each block is stored verbatim while the remaining gaps
// are stored as uint16 deltas.
func EncodeDelta(ids []uint32, blockSize int) []Block {
	if blockSize < 1 {
		blockSize = 64
	}
	blocks := make([]Block, 0, (len(ids)+blockSize-1)/blockSize)
	for start := 0; start < len(ids); start += blockSize {
		end := start + blockSize
		if end > len(ids) {
			end = len(ids)
		}
		block := Block{Base: ids[start]}
		for i := start + 1; i < end; i++ {
			block.Deltas = append(block.Deltas, uint16(ids[i]-ids[i-1]))
		}
		blocks = append(blocks, block)
	}
	return blocks
}

// Decode materialises the identifiers of a delta block.
func (b Block) Decode() []uint32 {
	out := make([]uint32, 0, len(b.Deltas)+1)
	current := b.Base
	out = append(out, current)
	for _, delta := range b.Deltas {
		current += uint32(delta)
		out = append(out, current)
	}
	return out
}

// Last returns the highest identifier contained in the block.
func (b Block) Last() uint32 {
	last := b.Base
	for _, delta := range b.Deltas {
		last += uint32(delta)
	}
	return last
}

// SkipSearch returns the index of the first block whose range can intersect
// the closed interval [lo, hi].
func SkipSearch(blocks []Block, lo, hi uint32) int {
	i := sort.Search(len(blocks), func(i int) bool {
		return blocks[i].Last() >= lo
	})
	if i < len(blocks) && blocks[i].Base <= hi {
		return i
	}
	if i >= len(blocks) {
		return len(blocks)
	}
	return i
}
