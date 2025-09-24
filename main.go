package main

import (
	"Tugas_1/input"
	"Tugas_1/service"
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("=============================================")
		fmt.Println("Selamat datang di aplikasi manajemen kontak!")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Tambah Kontak")
		fmt.Println("2. Lihat Kontak")
		fmt.Println("3. Update Kontak")
		fmt.Println("4. Hapus Kontak")
		fmt.Println("5. Keluar")
		fmt.Println("=============================================")
		inputUser := input.InputUser("Masukan pilihan menu: ")

		var numOpsiSelect int
		var name, phone, email, address string

		switch inputUser {
		case "1":
			name = input.InputUser("Masukan Nama: ")
			phone = input.InputUser("Masukan no Telp: ")
			email = input.InputUser("Masukan Email: ")
			address = input.InputUser("Masukan Alamat: ")
			fmt.Println()
			service.CreateContact(name, phone, email, address)
		case "2":
			service.ReadContact()
		case "3":
			numOpsiSelect, _ = strconv.Atoi(input.InputUser("Pilih kontak yang akan di ubah: "))
			updOpsiSelect, _ := strconv.Atoi(input.InputUser("Pilih data yang akan di ubah: "))
			fmt.Println("1. Nama")
			fmt.Println("2. No Telp")
			fmt.Println("3. Email")
			fmt.Println("4. Alamat")
			updValue := input.InputUser("Masukan data baru: ")
			service.UpdateContact(numOpsiSelect, updOpsiSelect, updValue)
		case "4":
			numOpsiSelect, _ = strconv.Atoi(input.InputUser("Masukan nomor kontak yang ingin dihapus: "))
			service.DeleteContact(numOpsiSelect)
		case "5":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
