package main

import (
	"fmt"
)

const (
	TarifDasar      = 5000
	TarifNextPerJam = 2000
	BiayaHariLibur  = 3000
)

func tarifParkir(durasi int, isMember bool, isHariLibur bool) float32 {
	// Validasi durasi parkir jika durasi kurang dari 0 jam
	if durasi <= 0 {
		return 0
	}

	var tarif float32

	// Logic Tarif dasar untuk 2 jam pertama
	if durasi <= 2 {
		tarif = TarifDasar
	} else { // logic tarif untuk 2 jam berikutnya
		nextJam := durasi - 2
		nextTarif := float32(nextJam) * TarifNextPerJam
		tarif = TarifDasar + nextTarif
	}

	// Menerapkan diskon untuk member
	if isMember {
		var disc float32
		if durasi <= 5 {
			disc = 0.50 // 50%
		} else {
			disc = 0.30 // 30%
		}

		tarif -= tarif * disc
	}

	// Menambahkan biaya pada hari libur
	if isHariLibur {
		tarif += BiayaHariLibur
	}

	return tarif
}

func main() {
	fmt.Println("==== Sistem Tarif Parkir ====")

	// Kasus 1
	// Seseorang parkir selama 4 jam, bukan member, dan bukan hari libur. Berapa tarif parkir totalnya?
	durasi, isMember, isHariLibur := 4, false, false
	tarif := tarifParkir(durasi, isMember, isHariLibur)
	fmt.Println("=== Kasus 1 ===")
	fmt.Println("Durasi Parkir		:", durasi)
	fmt.Println("Member 			:", isMember)
	fmt.Println("Hari Libur 		:", isHariLibur)
	fmt.Println("Total Biaya Parkir	: Rp.", tarif)

	fmt.Println("---------------------------")

	// Kasus 2
	// Seseorang parkir selama 2 jam, member, tetapi hari libur. Berapa tarif parkir totalnya?
	durasi2, isMember2, isHariLibur2 := 2, true, true
	tarif2 := tarifParkir(durasi2, isMember2, isHariLibur2)
	fmt.Println("=== Kasus 2 ===")
	fmt.Println("Durasi Parkir		:", durasi2)
	fmt.Println("Member 			:", isMember2)
	fmt.Println("Hari Libur 		:", isHariLibur2)
	fmt.Println("Total Biaya Parkir	: Rp.", tarif2)
}
