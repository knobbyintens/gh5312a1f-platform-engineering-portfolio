# delta-bitmap-index

Delta-compressed bitmap indexes for low-latency analytical queries.

`delta-bitmap-index` is a small Go library that stores sorted posting lists as
block-wise delta sequences. Each block keeps a base identifier plus fixed-size
deltas, and an in-memory block index lets conjunction queries skip blocks that
cannot intersect the query window.

## Features

- delta-encoded posting lists with configurable block size
- per-block skip pointers for selective conjunction
- roaring-style dense word intersection as a fallback
- bounded decode work proportional to matching blocks

## Layout

- `src/delta.go` — delta block encoding, decoding and skip search
- `src/bitmap.go` — bitmap container and intersection primitives
- `src/bitmap_test.go` — unit tests

## Usage

```go
idx := NewBitmapIndex(64)
idx.Add(42)
hits := idx.Intersect(10, 100)
```

## Status

Reference implementation used by the analytics engine experiments.
