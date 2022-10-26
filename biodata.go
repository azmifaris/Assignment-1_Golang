package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	var students []Biodata = []Biodata{
		{
			Nama: "Azmi Faris", Alamat: "Kebumen, Jawa Tengah", Pekerjaan: "Freelance", Alasan: "Ingin mempelajari Golang",
		},
		{
			Nama: "Wildan", Alamat: "Magelang, Jawa Tengah", Pekerjaan: "Developer", Alasan: "Ingin mempelajari lebih dalam Tentang Go",
		},
		{
			Nama: "Affan", Alamat: "Banjarnegara, Jawa Tengah", Pekerjaan: "Freelance", Alasan: "Tertarik Menggunakan bahasa Golang",
		},
		{
			Nama: "Faris", Alamat: "Sleman, Yogyakarta", Pekerjaan: "Front End Developer", Alasan: "Ingin mempelajari Bahasa Go",
		},
		{
			Nama: "Mufid", Alamat: "Kebumen, Jawa Tengah", Pekerjaan: "Fullstack Developer", Alasan: "Tertarik Menggunakan Golang",
		},
	}
	if len(os.Args) > 1 {
		index, _ := strconv.Atoi(os.Args[1])
		validasiAbsen(index, students)
	} else {
		fmt.Println("Nomor absen tidak ada")
	}
}

func tampil_data(index int, students []Biodata) {
	fmt.Println("No. Absen: ", index)
	fmt.Println("Nama: ", students[index-1].Nama)
	fmt.Println("Alamat: ", students[index-1].Alamat)
	fmt.Println("Pekerjaan: ", students[index-1].Pekerjaan)
	fmt.Println("Alasan: ", students[index-1].Alasan)
}
func validasiAbsen(index int, students []Biodata) {
	if index > 0 && index <= len(students) {
		tampil_data(index, students)
	} else {
		error_message(len(students))
	}
}
func error_message(length int) {
	fmt.Println("Peringatan!, absen hanya ada sampai absen ke- ", length)
}
