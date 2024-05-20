package main

import (
	"fmt"
)

func main() {
	var nama string = "Arif Laksonodhewo"
	var umur int8 = 19
	var marital_status bool = true

	fmt.Println("Nama ku adalah: ", nama)
	fmt.Println("Umur saya: ", umur)
	if(marital_status) {
		fmt.Println("Status saya saat ini sudah menikah")
	} else {
		fmt.Println("Status saya saat ini belum menikah")
	}

	fmt.Printf("Tipe data nama adalah %T" + "\n", nama)
	fmt.Printf("Tipe data umur adalah %T" + "\n", umur)
	fmt.Printf("Tipe data marital_status adalah %T", marital_status)
}