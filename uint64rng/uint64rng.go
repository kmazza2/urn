package uint64rng

// Any type with a parameter-free Next method returning uint64 can be used as a Uint64rng.
type Uint64rng interface {
	Next() uint64
}
