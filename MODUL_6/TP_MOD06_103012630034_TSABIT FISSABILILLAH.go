package main
import "fmt"

// func main(){
// 	var i, n, durasi_menit, total_menit, jam, menit int
// 	fmt.Scan(&n)
// 	total_menit = 0
// 	for i = 1; i <= n; i++ {
// 		fmt.Scan(&durasi_menit)
// 		total_menit = total_menit + durasi_menit
// 	}
// 	jam = total_menit / 60
// 	menit = total_menit % 60
// 	fmt.Printf("%d jam %d menit", jam, menit)
// }

func main() {
	var i, n, saldo_awal, kenaikan, saldo int
	fmt.Scan(&n, &saldo_awal, &kenaikan)
	saldo = saldo_awal
	for i = 1; i <= n; i++ {
		saldo = saldo + i*kenaikan
	}
	fmt.Println(saldo)
}
