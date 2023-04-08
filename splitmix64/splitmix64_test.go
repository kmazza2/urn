package splitmix64

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestMaxUint64(t *testing.T) {
	var go_max_uint64 uint64 = math.MaxUint64
	var gcc_max_unsigned_long uint64 = 18446744073709551615
	if go_max_uint64 != gcc_max_unsigned_long {
		t.Error(`Golang max uint64 not gcc max unsigned long`)
	}
}

func TestGeneratorAdvances(t *testing.T) {
	var rng SplitMix64 = NewSplitMix64(uint64(0))
	var first_number uint64 = rng.Next()
	var second_number uint64 = rng.Next()
	if first_number == second_number {
		t.Error(`RNG did not advance`)
	}
}

func TestAgainstCReference(t *testing.T) {
	var raw_ref_str_data []byte
	var err error
	raw_ref_str_data, err = ioutil.ReadFile(`testdata/c_impl_result`)
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
	for i := 0; i < 10; i++ {
		var rng SplitMix64 = NewSplitMix64(uint64(i))
		for j := 0; j < 10; j++ {
			data[10*i+j] = rng.Next()
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
		}
	}
	for k := 0; k < 100; k++ {
		if data[k] != ref_data[k] {
			t.Error(`Generated data does not match reference data`)
		}
	}
}
