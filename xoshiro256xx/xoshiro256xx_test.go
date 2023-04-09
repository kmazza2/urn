package xoshiro256xx

import (
	"github.com/kmazza2/urn/float64rng"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

// Make sure generator does not produce same value twice in a row.
// Note that xoshiro256xx actually *does* produce the same value twice in a row if seeded with math.MaxUint64 for all seeds. I am not sure if this is a flaw in the algorithm, or if instead it violates some principle of seeding the RNG. To be safe, the constructor panics if all seeds are math.MaxUint64.
func TestGeneratorAdvances(t *testing.T) {
	var rng Xoshiro256xx = NewXoshiro256xx(uint64(0), uint64(1), uint64(2), uint64(3))
	var first_number uint64 = rng.Next()
	var second_number uint64 = rng.Next()
	if first_number == second_number {
		t.Error(`RNG did not advance`)
	}
}

func TestAgainstCReference_uint64(t *testing.T) {
	var raw_ref_str_data []byte
	var err error
	raw_ref_str_data, err = ioutil.ReadFile(`testdata/uint64_from_splitmix64`)
	if err != nil {
		t.Error(`Could not open file containing test data`)
	}
	var ref_str_data []string = strings.Split(string(raw_ref_str_data), "\n")
	ref_str_data = ref_str_data[:len(ref_str_data)-1]
	var ref_data []uint64 = make([]uint64, 100)
	for i, str_num := range ref_str_data {
		var converted_num uint64
		var err error
		converted_num, err = strconv.ParseUint(str_num, 10, 64)
		if err != nil {
			t.Error(`Could not parse data`)
		}
		ref_data[i] = converted_num
	}
	var data []uint64 = make([]uint64, 100)
	var rng Xoshiro256xx = NewXoshiro256xx(
		16294208416658607535,
		7960286522194355700,
		487617019471545679,
		17909611376780542444)
	for i := 0; i < 100; i++ {
		data[i] = rng.Next()
	}
	for i := 0; i < 100; i++ {
		if data[i] != ref_data[i] {
			t.Error(`Generated data does not match reference data`)
		}
	}
}

// The maximum uint64 is 18446744073709551615. Hence if the random values are uniformly distributed the expectation is 18446744073709551615/2, which is approximately 9.2233720368547758075e18.
func TestExpectation(t *testing.T) {
	var expectation float64 = 9.2233720368547758075e18
	var mean float64 = 0.
	var samples = 100000
	var rng Xoshiro256xx = NewXoshiro256xx(
		16294208416658607535,
		7960286522194355700,
		487617019471545679,
		17909611376780542444)
	for i := 0; i < samples; i++ {
		mean += float64(rng.Next()) / float64(samples)
	}
	if math.Abs(expectation-mean) > 1e17 {
		t.Error(`Mean is far from expectation`)
	}
}

func TestPanicOnZeroSeed(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	var _ Xoshiro256xx = NewXoshiro256xx(0, 0, 0, 0)
	t.Error(`Did not panic when constructed with zero seed`)
}

func TestPanicOnMaxSeed(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	var _ Xoshiro256xx = NewXoshiro256xx(math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64)
	t.Error(`Did not panic when constructed with all seeds math.MaxUint64`)
}

// Make sure correct values are produced when wrapped with Float64rng.
func TestAgainstCReference_float64(t *testing.T) {
	var raw_ref_str_data []byte
	var err error
	raw_ref_str_data, err = ioutil.ReadFile(`testdata/float64_from_xoshiro256xx`)
	if err != nil {
		t.Error(`Could not open file containing test data`)
	}
	var ref_str_data []string = strings.Split(string(raw_ref_str_data), "\n")
	ref_str_data = ref_str_data[:len(ref_str_data)-1]
	var ref_data []float64 = make([]float64, 100)
	for i, str_num := range ref_str_data {
		var converted_num float64
		var err error
		converted_num, err = strconv.ParseFloat(str_num, 64)
		if err != nil {
			t.Error(`Could not parse data`)
		}
		ref_data[i] = converted_num
	}
	var data []float64 = make([]float64, 100)
	var src_rng Xoshiro256xx = NewXoshiro256xx(
		16294208416658607535,
		7960286522194355700,
		487617019471545679,
		17909611376780542444)
	var rng float64rng.Float64rng = float64rng.NewFloat64rng(&src_rng)
	for i := 0; i < 100; i++ {
		data[i] = rng.Next()
	}
	for i := 0; i < 100; i++ {
		if data[i] != ref_data[i] {
			t.Error(`Generated data does not match reference data`)
		}
	}
}
