package bitmapindex

import "testing"

func TestEncodeDecodeRoundTrip(t *testing.T) {
	ids := []uint32{1, 3, 4, 900, 901, 4000}
	blocks := EncodeDelta(ids, 2)
	if len(blocks) != 3 {
		t.Fatalf("block count = %d, want 3", len(blocks))
	}
	decoded := make([]uint32, 0)
	for _, b := range blocks {
		decoded = append(decoded, b.Decode()...)
	}
	if len(decoded) != len(ids) {
		t.Fatalf("round trip length = %d, want %d", len(decoded), len(ids))
	}
	for i := range ids {
		if decoded[i] != ids[i] {
			t.Fatalf("mismatch at %d: got %d want %d", i, decoded[i], ids[i])
		}
	}
}

func TestSkipSearch(t *testing.T) {
	blocks := EncodeDelta([]uint32{10, 11, 12, 500, 501}, 2)
	if got := SkipSearch(blocks, 500, 600); got != 2 {
		t.Fatalf("SkipSearch = %d, want 2", got)
	}
}

func TestIntersect(t *testing.T) {
	idx := NewBitmapIndex(4)
	idx.Add(5, 17, 42, 107)
	hits := idx.Intersect(10, 50)
	if len(hits) != 2 || hits[0] != 17 || hits[1] != 42 {
		t.Fatalf("unexpected hits: %v", hits)
	}
}
