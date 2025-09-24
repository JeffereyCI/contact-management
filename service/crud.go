package service

import (
	"Tugas_1/structure"
	"fmt"
)

// menggunakan angka dari 1
// var indexContact int = 1
var contactProject []structure.Contact

func CreateContact(Name string, Phone string, Email string, Address string) {
	contactProject = append(contactProject, structure.Contact{
		Name:    Name,
		Phone:   Phone,
		Email:   Email,
		Address: Address,
	})

	fmt.Println("Kontak berhasil ditambahkan!")
}

func ReadContact() {
	fmt.Println("=============================================")
	fmt.Println("Daftar Kontak:")

	for i, contact := range contactProject {
		fmt.Printf("%d. Nama: %s\n", i+1, contact.Name)
		fmt.Printf("   No Telp: %s\n", contact.Phone)
		fmt.Printf("   Email: %s\n", contact.Email)
		fmt.Printf("   Alamat: %s\n", contact.Address)
	}

	fmt.Println("=============================================")
}

func UpdateContact(numOpsi int, updOpsi int, updValue string) {
	numOpsi = numOpsi - 1
	switch updOpsi {
	case 1:
		contactProject[numOpsi].Name = updValue
	case 2:
		contactProject[numOpsi].Phone = updValue
	case 3:
		contactProject[numOpsi].Email = updValue
	case 4:
		contactProject[numOpsi].Address = updValue

		fmt.Println("Kontak berhasil di update!")
	}
}

func DeleteContact(numOpsi int) {
	var deleteID int
	for i := range contactProject {
		if i+1 == numOpsi {
			deleteID = i
			break
		}
	}
	contactProject = append(contactProject[:deleteID], contactProject[deleteID+1:]...)

	fmt.Println("Kontak berhasil dihapus!")
}
