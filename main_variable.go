package main

import "fmt"

func main () {
	var name string = "ahmad syahroni"
	var age float64 = 20
	var status bool = false
	ipk := 3.5
	var nickname string = "ahmad"
	var ptr *string = &name


	fmt.Println(
		"Nama saya adalah",
		name, "umur saya",
		age, "tahun",
		"status saya",
		status, "ipk saya",
		ipk, "nickname saya",
		nickname, "alamat pointer",
		ptr, "isi pointer",
		*ptr,
	)
}