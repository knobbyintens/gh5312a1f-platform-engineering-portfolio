package bitmapindex

import (
	"math/bits"
	"sort"
)

// BitmapIndex stores delta blocks for a single posting list.
type BitmapIndex struct {
	BlockSize int
	blocks    []Block
	dirty     bool
}

// NewBitmapIndex creates an index with the given block size.
func NewBitmapIndex(blockSize int) *BitmapIndex {
	if blockSize < 1 {
		blockSize = 64
	}
	return &BitmapIndex{BlockSize: blockSize, dirty: true}
}

// Add inserts all identifiers and re-encodes the posting list.
func (b *BitmapIndex) Add(ids ...uint32) {
	current := b.blocksToIDs()
	current = append(current, ids...)
	sort.Slice(current, func(i, j int) bool { return current[i] < current[j] })
	b.blocks = EncodeDelta(current, b.BlockSize)
	b.dirty = false
}

func (b *BitmapIndex) blocksToIDs() []uint32 {
	out := make([]uint32, 0)
	for _, block := range b.blocks {
		out = append(out, block.Decode()...)
	}
	return out
}

// Intersect returns identifiers in the closed interval [lo, hi]. The block
// index is used to skip blocks that cannot contribute.
func (b *BitmapIndex) Intersect(lo, hi uint32) []uint32 {
	start := SkipSearch(b.blocks, lo, hi)
	out := make([]uint32, 0)
	for i := start; i < len(b.blocks); i++ {
		block := b.blocks[i]
		if block.Base > hi {
			break
		}
		for _, id := range block.Decode() {
			if id >= lo && id <= hi {
				out = append(out, id)
			}
		}
	}
	return out
}

// WordPopcount is a small helper used by dense bitmap fallbacks.
func WordPopcount(word uint64) int {
	return bits.OnesCount64(word)
}
