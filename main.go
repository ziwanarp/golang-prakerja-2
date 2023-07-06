package main

import "fmt"

func bilanganPrima(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func kelipatan7(n float32) bool {
	result := n / 7
	return result == float32(int(result))
}

func luasTrapesium(alas, sisi, tinggi float32) float32 {
	var luas float32 = (alas + sisi) * tinggi / 2
	return luas
}

func main() {
	// input angka yang akan di seleksi
	angka := 701

	// ============= Menentukan bilangan prima ============== \\
	if bilanganPrima(angka) {
		fmt.Println(angka, "adalah bilangan prima.")
	} else {
		fmt.Println(angka, "bukan bilangan prima.")
	}

	// ============= Menentukan kelipatan 7 ============== \\
	input := float32(angka)
	if kelipatan7(input) {
		fmt.Println(angka, "adalah kelipatan 7.")
	} else {
		fmt.Println(angka, "bukan kelipatan 7.")
	}

	// ============= Menentukan Luas Trapesium ============== \\
	var alas float32 = 3
	var sisi float32 = 4
	var tinggi float32 = 3
	fmt.Println("Luas Trapesium : ", luasTrapesium(alas, sisi, tinggi))

}
