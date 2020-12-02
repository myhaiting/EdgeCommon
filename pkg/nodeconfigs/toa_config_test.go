package nodeconfigs

import (
	"runtime"
	"testing"
)

func TestTOAConfig_RandLocalPort(t *testing.T) {
	{
		toa := &TOAConfig{}
		err := toa.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(toa.RandLocalPort())
	}
	{
		toa := &TOAConfig{
			MinLocalPort: 1,
			MaxLocalPort: 2,
		}
		err := toa.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(toa.RandLocalPort())
	}
}

func BenchmarkTOAConfig_RandLocalPort(b *testing.B) {
	runtime.GOMAXPROCS(1)

	toa := &TOAConfig{
		MinLocalPort: 1,
		MaxLocalPort: 2,
	}
	_ = toa.Init()

	for i := 0; i < b.N; i++ {
		_ = toa.RandLocalPort()
	}
}