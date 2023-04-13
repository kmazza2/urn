package uint64tofloat64

import (
	u64r "github.com/kmazza2/urn/uint64rng"
)

// Uint64toFloat64 objects contain a uint64 RNG.
type Uint64toFloat64 struct {
	src u64r.Uint64rng
}

// Constructor for Uint64toFloat64 objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewUint64toFloat64(src u64r.Uint64rng) Uint64toFloat64 {
	return Uint64toFloat64{src}
}

// Generates the next pseudorandom value from the source uint64 generator, then converts it to a float64.
func (rng Uint64toFloat64) Next() float64 {
	var x uint64 = rng.src.Next()
	return float64(x>>11) * float64(0x1.0p-53)
}
