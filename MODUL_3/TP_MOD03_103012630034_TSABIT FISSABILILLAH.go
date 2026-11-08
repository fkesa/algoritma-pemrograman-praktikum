package main
import "fmt"

func konversi_waktu() {
	var total_detik, jam, menit, detik int

	fmt.Print("Masukkan total detik: ")
	fmt.Scan(&total_detik)
	// total_detik = 1456
	jam = total_detik / 3600
	menit = (total_detik % 3600) / 60
	detik = total_detik % 60
	fmt.Println(jam, menit, detik)
	fmt.Println(total_detik, "=", jam, "jam", menit, "menit", detik, "detik")
}

func wfc() {
	var angka, hari_awal, hari_akhir int
	var hari_arr []string

	// array ditambah string kosong diawal biar awal hari mulai dari 1
	hari_arr = []string{"","Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	hari_awal = 4 // Kamis
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	if angka >= 1 && angka <= 1000 {
		hari_akhir = ((hari_awal - 1 + angka) % 7) + 1
		fmt.Println(hari_akhir)
		fmt.Printf("%d hari setelah hari %s (%d) adalah %s (%d)\n",
			angka, hari_arr[hari_awal], hari_awal, hari_arr[hari_akhir], hari_akhir)
	} else {
		fmt.Println("Masukkan angka range 1-1000")
	}

}

func main() {
	konversi_waktu()
	konversi_waktu()
	konversi_waktu()
	wfc()
	wfc()
	wfc()
}
