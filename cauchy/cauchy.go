package cauchy

import (
	f64r "github.com/kmazza2/urn/float64rng"
	"math"
)

// Cauchyrng objects contain a float64 RNG, a float64 for alpha, and a float64 for beta.
type Cauchyrng struct {
	src   f64r.Float64rng
	alpha float64
	beta  float64
}

// Constructor for Cauchyrng objects. Intended to prevent client code from directly modifying RNG state after initialization.
func NewCauchyrng(src f64r.Float64rng, alpha float64, beta float64) Cauchyrng {
	if beta <= 0 {
		panic(`beta must be greater than 0`)
	}
	return Cauchyrng{src, alpha, beta}
}

// Generates the next pseudorandom value.
func (rng Cauchyrng) Next() float64 {
	var u float64
	u = rng.src.Next()
	return rng.alpha + rng.beta*math.Tan(math.Pi*(u-0.5))
}
