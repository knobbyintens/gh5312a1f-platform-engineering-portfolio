package router

// Pool describes a query engine pool and its current load.
type Pool struct {
	Name     string
	Capacity int
	InFlight int
}

// Choose returns the pool with the largest remaining capacity.
func Choose(pools []Pool) string {
	best := ""
	bestRemaining := -1
	for _, p := range pools {
		remaining := p.Capacity - p.InFlight
		if remaining > bestRemaining {
			best = p.Name
			bestRemaining = remaining
		}
	}
	return best
}
