package main
import "fmt"

func main(){
	var panjang_sisi_kotak, keliling_pita, panjang_simpul, total_panjang_pita int
	panjang_simpul = 20
	fmt.Printf("Masukkan panjang_sisi_kotak = ")
	fmt.Scan(&panjang_sisi_kotak)
	keliling_pita = 4 * panjang_sisi_kotak
	fmt.Println("Pita Keliling = ", keliling_pita)
	total_panjang_pita = keliling_pita + panjang_simpul
	fmt.Println("Pita Total = ", total_panjang_pita)
}
