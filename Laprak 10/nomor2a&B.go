func galiKartu(d *Dominoes, kartu Domino) {

	var ambil Domino
	ketemu := false

	for d.jumlah > 0 && !ketemu {

		ambil = ambilKartu(d)

		if ambil.sisi1 == kartu.sisi1 ||
			ambil.sisi2 == kartu.sisi1 ||
			ambil.sisi1 == kartu.sisi2 ||
			ambil.sisi2 == kartu.sisi2 {

			fmt.Println("Kartu cocok ditemukan:")
			fmt.Println(ambil)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada kartu yang cocok")
	}
}

func sepasangKartu(a Domino, b Domino) bool {

	total := nilaiKartu(a) + nilaiKartu(b)

	return total == 12
}