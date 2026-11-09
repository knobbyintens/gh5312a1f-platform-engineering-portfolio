"""Serve calibrated scan-cost predictions."""

import json


def load_coefficients(path):
    with open(path, encoding="utf-8") as fh:
        return json.load(fh)


def predict(coefficients, rows, selectivity):
    base = coefficients["base"]
    per_row = coefficients["per_row"]
    return base + per_row * rows * (1.0 - selectivity)
