// for - range jika diterapkan pada channel berfungsi untuk handle penerimaan data.
// Setiap kali ada pengiriman data via channel, maka akan men-trigger perulangan for - range.
// Perulangan akan berlangsung terus-menerus seiring pengiriman data ke channel yang dipergunakan.
// Dan perulangan hanya akan berhenti jika channel yang digunakan tersebut di close atau di-non-aktifkan.
// Fungsi close digunakan utuk me-non-aktifkan channel.

package main

import (
	"fmt"
	"runtime"
)

//parameter hanya bisa di gunakan untuk mengirim data
func SendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		data := fmt.Sprintf("data : %d", i)
		ch <- data
	}
	close(ch)
}

// parameter hanya bisa di gunakan untuk menerima data
func PrintMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	go SendMessage(message)
	PrintMessage(message)

}
