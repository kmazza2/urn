package cunif

import (
	f64r "github.com/kmazza2/urn/float64rng"
)

// Cunifrng objects contain a float64 RNG, a lower limit l, and an upper limit u.
type Cunifrng struct {
	src f64r.Float64rng
	l   float64
	u   float64
}

// Constructor for Cunifrng objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewCunifrng(src f64r.Float64rng, l float64, u float64) Cunifrng {
	if l > u {
		panic(`l must be less than or equal to u`)
	}
	return Cunifrng{src, l, u}
}

// Generates the next pseudorandom value.
func (rng Cunifrng) Next() float64 {
	return rng.src.Next()*(rng.u-rng.l) + rng.l
}
