package main
import "fmt"

type mahasiswa struct {
	nama        string
	nim         string
	kuis        float64
	tugas       float64
	praktikum   float64
	uts         float64
	uas         float64
	nilai_akhir float64
}

// SOAL 1 NILAI
// func main() {
// 	var m mahasiswa
// 	fmt.Scan(&m.nama, &m.nim, &m.tugas, &m.kuis, &m.praktikum, &m.uts, &m.uas)
// 	fmt.Println(m.nim)
// 	fmt.Println(m.nama)
// 	fmt.Println(m.tugas)
// 	fmt.Println(m.praktikum)
// 	fmt.Println(m.kuis)
// 	fmt.Println(m.uts)
// 	fmt.Println(m.uas)
// 	m.nilai_akhir = 15.0/100.0*m.tugas + 15.0/100.0*m.kuis + 20.0/100.0*m.praktikum + 25.0/100.0*m.uts + 25.0/100.0*m.uas
// 	fmt.Println(m.nilai_akhir)
// }

// SOAL 2 SENSOR
func main() {
	var a, b, c int
	var sensor_active bool
	fmt.Scan(&a, &b, &c)
	sensor_active = (a == 1 && b == 1) || (b == 1 && c == 1) || (c == 1 && a == 1)
	fmt.Println(sensor_active)
}
