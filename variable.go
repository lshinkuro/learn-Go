package main

import (
	"fmt"
)

var nama = "ahmad"

func main() {
	//cara membuat pointer
	var address = "bandung"
	var address1 *string = &address

	fmt.Println("alamat bukan pointer :", address)
	/*jika ingin menampilkan isi pointer gunakan *namapointer
	jika nilai pointer di rubah maka akan berefek kenilai lainya juga */

	fmt.Println("alamat pointer :", *address1)

	*address1 = "Sumatra Barat"

	fmt.Println("alamat setelah mengubah Pointer", address)

	var value int = 4
	fmt.Println("before value change", value)

	Change(&value, 30)

	fmt.Println("after change Pointer", value)

}

//ini adalah cara membuat pointer di dalam parameter

func Change(original *int, value int) {
	*original = value
}
