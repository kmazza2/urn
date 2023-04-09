package xoshiro256xx

import (
	"math"
)

// Xoshiro256xx objects contain the state of a Xoshiro256xx RNG.
type Xoshiro256xx struct {
	state0 uint64
	state1 uint64
	state2 uint64
	state3 uint64
}

// Creates a new Xoshiro256xx object (which happens to be determined entirely by its state, which on initialization is equal to its seed).
func NewXoshiro256xx(
	seed0 uint64,
	seed1 uint64,
	seed2 uint64,
	seed3 uint64) Xoshiro256xx {
	switch {
	case seed0 == math.MaxUint64 && seed1 == seed0 && seed2 == seed0 && seed3 == seed0:
		panic(`Setting all Xoshiro256xx seeds to math.MaxUint64 results in two in a row.`)
	case seed0 == 0 && seed1 == 0 && seed2 == 0 && seed3 == 0:
		panic(`At least one Xoshiro256xx seed must be nonzero`)
	default:
		return Xoshiro256xx{seed0, seed1, seed2, seed3}
	}
}

func rotl(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}

// Returns the next pseudorandom value produced by the xoshiro256xx algorithm given the state in rng.
func (rng *Xoshiro256xx) Next() uint64 {
	var result uint64 = rotl(rng.state1*5, 7) * 9
	var t uint64 = rng.state1 << 17
	rng.state2 ^= rng.state0
	rng.state3 ^= rng.state1
	rng.state1 ^= rng.state2
	rng.state0 ^= rng.state3
	rng.state2 ^= t
	rng.state3 = rotl(rng.state3, 45)
	return result
}
