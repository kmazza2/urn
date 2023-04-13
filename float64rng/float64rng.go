package float64rng

// Any type with a parameter-free Next method returning float64 can be used as a Float64rng.
type Float64rng interface {
	Next() float64
}
