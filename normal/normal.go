package normal

import (
	f64r "github.com/kmazza2/urn/float64rng"
	"math"
)

// Normalrng objects contain a float64 RNG, a bool indicating if the next random number produced has been cached, a float64 holding cached random numbers, a float64 for the variance, and a float64 for the mean.
type Normalrng struct {
	src      f64r.Float64rng
	cached   bool
	cache    float64
	mean     float64
	variance float64
}

// Constructor for Normalrng objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewNormalrng(src f64r.Float64rng, mean float64, variance float64) Normalrng {
	if variance <= 0 {
		panic(`variance must be greater than 0`)
	}
	return Normalrng{src, false, 0., mean, variance}
}

// Generates the next pseudorandom value.
func (rng Normalrng) Next() float64 {
	if rng.cached {
		rng.cached = false
		return rng.cache
	}
	var u1 float64
	var u2 float64
	u1 = rng.src.Next()
	u2 = rng.src.Next()
	rng.cache = rng.mean + rng.variance*math.Sqrt(-2.*math.Log(u1))*math.Sin(2*math.Pi*u2)
	rng.cached = true
	return rng.mean + rng.variance*math.Sqrt(-2.*math.Log(u1))*math.Cos(2*math.Pi*u2)
}
