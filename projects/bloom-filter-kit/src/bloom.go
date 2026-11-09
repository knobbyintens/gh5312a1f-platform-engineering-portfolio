package bloom

// Filter is a fixed-size bloom filter backed by a bit array.
type Filter struct {
	bits []uint64
	k    int
}

// New builds an empty filter sized for n items at the given false-positive rate.
func New(n int, k int) *Filter {
	return &Filter{bits: make([]uint64, n/64+1), k: k}
}

// Add marks the hashed positions for the value.
func (f *Filter) Add(value string) {
	for i := 0; i < f.k; i++ {
		h := hash(value, i) % uint64(len(f.bits)*64)
		f.bits[h/64] |= 1 << (h % 64)
	}
}

func hash(value string, seed int) uint64 {
	var h uint64 = 1469598103934665603
	for _, c := range value {
		h ^= uint64(c)
		h *= 1099511628211
	}
	return h + uint64(seed)*0x9E3779B97F4A7C15
}
