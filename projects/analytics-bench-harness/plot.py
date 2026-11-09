"""Plot benchmark results produced by bench.py."""

import csv

import matplotlib.pyplot as plt


def main():
    rows = []
    with open("results.csv", newline="") as fh:
        for row in csv.DictReader(fh):
            rows.append(row)
    fig, ax = plt.subplots()
    ax.plot([int(r["cardinality"]) for r in rows], [float(r["seconds"]) for r in rows])
    ax.set_xlabel("cardinality")
    ax.set_ylabel("seconds")
    fig.savefig("latency.png")


if __name__ == "__main__":
    main()
