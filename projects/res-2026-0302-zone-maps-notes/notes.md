# Notes: Columnar Zone Maps (RES-2026-0302)

Zone maps summarise min/max values per stored column chunk so a scan can skip
chunks that cannot satisfy the predicate. The note set records the reported
pruning ratios and the open question of update maintenance. No code accompanies
this report.
