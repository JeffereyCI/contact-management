package service

import (
	"Tugas_1/structure"
	"fmt"
)

var contactProject []structure.Contact

func CreateContact(Name string, Phone string, Email string, Address string) {
	fmt.Println("Kontak berhasil ditambahkan!")
}

func ReadContact() {
	fmt.Println("=============================================")
	fmt.Println("Daftar Kontak:")

	// loop

	fmt.Println("=============================================")
}

func UpdateContact(numOpsi string) {

	fmt.Println("Kontak berhasil di update!")
}

func DeleteContact(numOpsi string) {

	fmt.Println("Kontak berhasil dihapus!")
}
