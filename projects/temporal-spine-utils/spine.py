"""Align event partitions to a canonical temporal spine."""

from dataclasses import dataclass


@dataclass
class Partition:
    name: str
    start_hour: int
    end_hour: int


def align(partitions, spine_start, spine_end):
    """Return the partitions clipped to the shared spine window."""
    out = []
    for p in partitions:
        start = max(p.start_hour, spine_start)
        end = min(p.end_hour, spine_end)
        if start < end:
            out.append(Partition(p.name, start, end))
    return out
