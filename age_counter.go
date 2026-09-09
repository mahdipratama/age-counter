package agecounter

type Age struct {
	Tanggal int
	Bulan   int
	Tahun   int
}

func AgeCounter(birtdate Age, today Age) (Age, error) {

	calculatedAge := Age{}
	day := 0
	bulan := 0
	tahun := 0

	// Calculate days
	if today.Tanggal < birtdate.Tanggal {
		today.Bulan -= 1
		today.Tanggal += 30
	}
	day = today.Tanggal - birtdate.Tanggal

	// Calculate months
	if today.Bulan < birtdate.Bulan {
		today.Tahun -= 1
		today.Bulan += 12
	}
	bulan = today.Bulan - birtdate.Bulan

	// Calculate years
	tahun = today.Tahun - birtdate.Tahun

	calculatedAge = Age{
		day,
		bulan,
		tahun,
	}

	return calculatedAge, nil
}
