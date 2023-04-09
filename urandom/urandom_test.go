package urandom

import (
	"testing"
	"math"
)

// Make sure the generator doesn't produce the same value twice in a row.
func TestGeneratorAdvances(t *testing.T) {
	var rng Urandom = NewUrandom()
	var first_number uint64 = rng.Next()
	var second_number uint64 = rng.Next()
	if first_number == second_number {
		t.Error(`RNG did not advance`)
	}
}

// The maximum uint64 is 18446744073709551615. Hence if the random values are uniformly distributed the expectation is 18446744073709551615/2, which is approximately 9.2233720368547758075e18.
func TestExpectation(t *testing.T) {
	var expectation float64 = 9.2233720368547758075e18
	var mean float64 = 0.
	var samples = 100000
	var rng Urandom = NewUrandom()
	for i := 0; i < samples; i++ {
		mean += float64(rng.Next()) / float64(samples)
	}
	if math.Abs(expectation - mean) > 1e17 {
		t.Error(`Mean is far from expectation`)
	}
}
