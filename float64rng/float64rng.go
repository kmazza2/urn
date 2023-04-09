package float64rng

// Any type with a parameter-free Next method returning uint64 can be used as a source generator for Float64rng.
type Uint64rng interface {
	Next() uint64
}

// Float64rng objects contain a uint64 RNG.
type Float64rng struct {
	src Uint64rng
}

// Constructor for Float64rng objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewFloat64rng(src Uint64rng) Float64rng {
	return Float64rng{src}
}

// Generates the next pseudorandom value from the source uint64 generator, then converts it to a float64.
func (rng Float64rng) Next() float64 {
	var x uint64 = rng.src.Next()
	return float64(x>>11) * float64(0x1.0p-53)
}
