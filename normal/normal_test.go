package normal

import (
	f64r "github.com/kmazza2/urn/float64rng"
	x256xx "github.com/kmazza2/urn/xoshiro256xx"
	"testing"
)

func TestGenerator(t *testing.T) {
	var src_rng x256xx.Xoshiro256xx = x256xx.NewXoshiro256xx(
		16294208416658607535,
		7960286522194355700,
		487617019471545679,
		17909611376780542444)
	var rng f64r.Float64rng = f64r.NewFloat64rng(&src_rng)
	var normal_rng = NewNormalrng(rng, 2.5, -9.2)
	for i := 0; i < 10000; i++ {
		_ = normal_rng.Next()
	}
	t.Error(`Not implemented`)
}
