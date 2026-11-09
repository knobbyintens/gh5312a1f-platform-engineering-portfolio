"""Benchmark third-party analytics index packages.

Note: this harness does not contain an index implementation. It imports the
external ``delta_bitmap_reference`` package, runs selective conjunctions on
synthetic identifiers, and writes results to ``results.csv``.
"""

import csv
import time

import delta_bitmap_reference as reference


def run_case(cardinality, queries, repeat=5):
    index = reference.build_index(cardinality)
    best = float("inf")
    for _ in range(repeat):
        start = time.perf_counter()
        reference.run_queries(index, queries)
        best = min(best, time.perf_counter() - start)
    return cardinality, queries, best


def main():
    cases = [(100_000, 64), (1_000_000, 64), (4_000_000, 16)]
    with open("results.csv", "w", newline="") as fh:
        writer = csv.writer(fh)
        writer.writerow(["cardinality", "queries", "seconds"])
        for case in cases:
            writer.writerow(run_case(*case))


if __name__ == "__main__":
    main()
