package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus             Kubus   = Kubus{4}
	volumeSeharusnya  float64 = 64
	luasSeharusnya    float64 = 96
	kelilingSeharunya float64 = 48
)

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Salah! seharunya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("Salah! seharunysa %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kelilingSeharunya)

	if kubus.Keliling() != kelilingSeharunya {
		t.Errorf("Salah! seharusnya %.2f", kelilingSeharunya)
	}
}

func TestAssertHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestAssertHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharunya, "perhitungan keliling salah")
}

func TestAssertHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}
