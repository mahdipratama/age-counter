package agecounter_test

import (
	"agecounter"
	"testing"
)

func TestAgeCounter_ReturnsCorrectCalculatedAge(t *testing.T) {
	t.Parallel()

	today_test := agecounter.Age{
		Tanggal: 9,
		Bulan:   9,
		Tahun:   2026,
	}

	want := agecounter.Age{
		Tanggal: 28,
		Bulan:   3,
		Tahun:   28,
	}

	got, err := agecounter.AgeCounter(
		agecounter.Age{
			Tanggal: 11,
			Bulan:   5,
			Tahun:   1998,
		}, today_test)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want: %#v, got %#v", want, got)
	}

}
