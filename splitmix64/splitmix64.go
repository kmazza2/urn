package splitmix64

// SplitMix64 objects contain the state of a SplitMix64 RNG.
type SplitMix64 struct {
	state uint64
}

// Creates a new SplitMix64 object (which happens to be determined entirely by its state, which on initialization is equal to its seed).
func NewSplitMix64(seed uint64) SplitMix64 {
	var rng SplitMix64 = SplitMix64{seed}
	return rng
}

// Returns the next pseudorandom value produced by the SplitMix64 algorithm given the state in rng.
func (rng *SplitMix64) Next() uint64 {
	rng.state += 0x9e3779b97f4a7c15
	var z uint64 = rng.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
