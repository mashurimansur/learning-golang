package unit_test_test

import (
	unit_test "github.com/mashurimansur/learning-golang/unit-test"
	"testing"
)

var (
	kubus              unit_test.Kubus = unit_test.Kubus{4}
	volumeSeharusnya   float64         = 64
	luasSeharusnya     float64         = 96
	kelilingSeharusnya float64         = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Valume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Salah! harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("Salah! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
	}
}

func BenchmarkKubus_Luas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
