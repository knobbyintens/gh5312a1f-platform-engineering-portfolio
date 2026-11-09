# Notes: Delta-Compressed Bitmap Indexes (RES-2026-0417)

**Report:** RES-2026-0417 — "Delta-Compressed Bitmap Indexes for Low-Latency Analytics"

## One-paragraph summary

Posting lists are sorted, so identifiers can be stored as block-wise delta
sequences. Each block carries a base identifier and fixed-size deltas. Queries
binary-search the block index, skip blocks outside the query range, and decode
only candidate blocks. The reported evaluation claims 2.3x faster selective
conjunctions than a Roaring bitmap baseline at roughly equal index size.

## Points worth checking in review

- The evaluation uses a synthetic append-only table alongside TPC-H; it is not a
  production query trace.
- Delta blocks are fixed-size; no discussion of skewed block occupancy.
- Memory overhead of the block index is excluded from the size figures.

## Open questions

1. How does block size interact with cache lines on the target CPU?
2. Are updates (deletes/tombstones) supported without full rebuild?
3. Which runtime should host a port if we adopt the method?

## Status

Notes only — no implementation. Other projects may contain related code;
this project does not.
