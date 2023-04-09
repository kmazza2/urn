package dunif

// Any type with a parameter-free Next method returning float64 (in (0, 1)) can be used as a source generator for Dunif. The numbers returned by Next need to be approximately Uniform(0, 1).
type Float64rng interface {
	Next() float64
}

// Dunifrng objects contain a float64 RNG and a value n.
type Dunifrng struct {
	src Float64rng
	n   uint64
}

// Constructor for Float64rng objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewDunifrng(src Float64rng, n uint64) Dunifrng {
	if n == 0 {
		panic(`n must be greater than or equal to 1`)
	}
	return Dunifrng{src, n}
}

// Generates the next pseudorandom value.
func (rng Dunifrng) Next() uint64 {
	var mass float64 = 1. / float64(rng.n)
	var u float64 = rng.src.Next()
	var i uint64
	for i = 1; float64(i)*mass <= u; i++ {
	}
	return i
}
