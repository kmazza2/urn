package urandom

import (
	"encoding/binary"
	"io"
	"os"
)

// Urandom is an empty struct, provided only for consistency with other RNG packages in this module.
type Urandom struct {
}

// Creates a new Urandom object.
func NewUrandom() Urandom {
	var rng Urandom = Urandom{}
	return rng
}

// Returns the next pseudorandom value provided by the OS, or panics if the OS fails to provide a value.
func (rng *Urandom) Next() uint64 {
	var r *os.File
	var err error
	r, err = os.Open("/dev/urandom")
	if err != nil {
		panic(err)
	}
	defer func() {
		var err error
		err = r.Close()
		if err != nil {
			panic(err)
		}
	}()
	var lr = io.LimitReader(r, 8)
	var bytes []byte = make([]byte, 8)
	n, err := lr.Read(bytes)
	if n != 8 {
		panic(`Failed to read from /dev/urandom`)
	}
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(bytes)
}
