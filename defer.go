// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
// Sedangkan Exit digunakan untuk menghentikan program secara paksa
// (ingat, menghentikan program, tidak seperti return yang hanya menghentikan blok kode).

package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("hallo")

	func() {
		fmt.Println("halo semuanyaa")
		defer fmt.Println("hallo")
	}()

	fmt.Println("halo semuanyaa")
}
