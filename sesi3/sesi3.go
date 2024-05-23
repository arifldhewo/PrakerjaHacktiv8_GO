package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type person struct {
		Nama      string
		Alamat    string
		Pekerjaan string
		Alasan    string
	}

	var people []person

	people = append(people,
		person{
			Nama: "nil",
			Alamat: "nil",
			Pekerjaan: "nil",
			Alasan: "nil",
		},
		person{
			Nama:      "Morgott the Omen King",
			Alamat:    "Jl Jend A Yani 1037, Jawa Barat",
			Pekerjaan: "Guardian",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
		person{
			Nama:      "Queen Marika",
			Alamat:    "Jl. Manggarai Utara I No.1 Jakarta selatan, Dki Jakarta",
			Pekerjaan: "the Eternal",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
		person{
			Nama:      "Melina",
			Alamat:    "Jl Wijaya II Wijaya Graha Puri Bl B/4, Dki Jakarta",
			Pekerjaan: "the Messenger",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
		person{
			Nama:      "Malenia",
			Alamat:    "Jl Mesjid 71-AB, Sumatera Utara",
			Pekerjaan: "Blade of Miquella",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
		person{
			Nama:      "Ranni",
			Alamat:    "Jl Ibnu Taimia IV Kompl IAIN, Dki Jakarta",
			Pekerjaan: "the Witch",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
		person{
			Nama:      "Fia",
			Alamat:    "Jl Pelabuhan Kalibaru 15, Jakarta",
			Pekerjaan: "the Deathbed Companion",
			Alasan:    "Improving Skills and Adding new knowledge",
		},
	)

	if len(os.Args) > 1  {

		valConv, errConvOs := strconv.Atoi(os.Args[1])
		if valConv < len(people) {

			if (errConvOs != nil) {
				fmt.Println("Tidak menerima input selain angka")
			} else {
				fmt.Println("Nama : ", people[valConv].Nama)
				fmt.Println("Alamat : ", people[valConv].Alamat)
				fmt.Println("Pekerjaan : ", people[valConv].Pekerjaan)
				fmt.Println("Alasan: ", people[valConv].Alasan)
			}

		} else {

			fmt.Print("Data tidak ditemukan")
		}
	} else {
		
		fmt.Print("Anda tidak memberikan data apapun!!")
	}
}