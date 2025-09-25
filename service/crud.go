package service

import (
	"Tugas_1/structure"
	"fmt"
)

var ContactProject []structure.Contact

func CreateContact(Name string, Phone string, Email string, Address string) {
	ContactProject = append(ContactProject, structure.Contact{
		Name:    Name,
		Phone:   Phone,
		Email:   Email,
		Address: Address,
	})

	fmt.Println("Kontak berhasil ditambahkan!")
}

func ReadContact() {
	fmt.Println("")
	fmt.Println("=============================================")
	fmt.Println("Daftar Kontak:")

	if len(ContactProject) == 0 {
		fmt.Println("		Kontak kosong")
		fmt.Println("  silahkan tambahkan kontak terlebih dahulu")
		fmt.Println("=============================================")
		fmt.Println()
		return
	} else if len(ContactProject) > 0 {
		fmt.Println("---------------------------------------------")
		for i, contact := range ContactProject {
			fmt.Printf("%d. Nama : %s\n", i+1, contact.Name)
			fmt.Printf("No Telp	: %s\n", contact.Phone)
			fmt.Printf("Email	: %s\n", contact.Email)
			fmt.Printf("Alamat	: %s\n", contact.Address)
			fmt.Println("---------------------------------------------")
		}
	}

	fmt.Println("=============================================")
	fmt.Println()
}

func UpdateContact(numOpsi int, updOpsi int, updValue string) {
	numOpsi = numOpsi - 1

	if len(ContactProject) == 0 {
		fmt.Println()
		fmt.Println("Kontak kosong, silahkan tambahkan kontak terlebih dahulu")
		fmt.Println()
		return
	} else if len(ContactProject) > 0 && (numOpsi <= len(ContactProject)) {
		switch updOpsi {
		case 1:
			ContactProject[numOpsi].Name = updValue
		case 2:
			ContactProject[numOpsi].Phone = updValue
		case 3:
			ContactProject[numOpsi].Email = updValue
		case 4:
			ContactProject[numOpsi].Address = updValue

			fmt.Println("Kontak berhasil di update!")
			fmt.Println()
		}
	} else {
		fmt.Println("kontak yang dipilih tidak ada")
		fmt.Println()
	}
}

func DeleteContact(numOpsi int) {
	// var deleteID int
	// for i := range ContactProject {
	// 	if i+1 == numOpsi {
	// 		deleteID = i
	// 		break
	// 	}
	// }
	// ContactProject = append(ContactProject[:deleteID], ContactProject[deleteID+1:]...)

	if len(ContactProject) == 0 {
		fmt.Println("Kontak kosong, silahkan tambahkan kontak terlebih dahulu")
		fmt.Println()
		return
	} else if len(ContactProject) > 0 && (numOpsi <= len(ContactProject)) {
		numOpsi = numOpsi - 1
		ContactProject = append(ContactProject[:numOpsi], ContactProject[numOpsi+1:]...)

		fmt.Println("Kontak berhasil dihapus!")
		fmt.Println()
	} else {
		fmt.Println("kontak yang dipilih tidak ada")
		fmt.Println()
	}
}
