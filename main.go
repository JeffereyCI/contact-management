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
		fmt.Println("=============================================")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Tambah Kontak")
		fmt.Println("2. Lihat Kontak")
		fmt.Println("3. Update Kontak")
		fmt.Println("4. Hapus Kontak")
		fmt.Println()
		fmt.Println("---------------------------------------------")
		fmt.Println("q. Keluar")
		fmt.Println("=============================================")
		fmt.Println()

		inputUser := input.InputUser("Masukan pilihan menu: ")

		switch inputUser {
		case "1":
			fmt.Println()
			fmt.Println("----------------------------------------------")
			fmt.Println("q. untuk kembali ke menu utama")
			fmt.Println("----------------------------------------------")

			name := input.InputUser("Masukan Nama   : ")
			if name == "q" {
				fmt.Println()
				continue
			}
			phone := input.InputUser("Masukan No Telp: ")
			if phone == "q" {
				fmt.Println()
				continue
			}
			email := input.InputUser("Masukan Email  : ")
			if email == "q" {
				fmt.Println()
				continue
			}
			address := input.InputUser("Masukan Alamat : ")
			if address == "q" {
				fmt.Println()
				continue
			}

			fmt.Println("----------------------------------------------")
			service.CreateContact(name, phone, email, address)
			fmt.Println("----------------------------------------------")
			fmt.Println()

		case "2":
			service.ReadContact()

		case "3":
			service.ReadContact()
			fmt.Println("----------------------------------------------")
			fmt.Println("pilih q untuk kembali ke menu utama")
			fmt.Println("----------------------------------------------")

			numStr := input.InputUser("Pilih kontak yang akan diubah: ")
			fmt.Println()
			if numStr == "q" {
				fmt.Println()
				continue
			}
			numOpsiSelect, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Input tidak valid.")
				fmt.Println()
				continue
			}

			if len(service.ContactProject) == 0 {
				fmt.Println("kontak yang dipilih tidak ada")
				fmt.Println()
				continue
			} else if len(service.ContactProject) > 0 && (numOpsiSelect <= len(service.ContactProject)) {
				fmt.Println()
				fmt.Println("==============================================")
				fmt.Println("1. Nama")
				fmt.Println("2. No Telp")
				fmt.Println("3. Email")
				fmt.Println("4. Alamat")
				fmt.Println("==============================================")
				fmt.Println()

				updStr := input.InputUser("Pilih bagian yang akan diubah: ")

				if updStr == "q" {
					fmt.Println()
					continue
				}
				updOpsiSelect, err := strconv.Atoi(updStr)
				if err != nil {
					fmt.Println("Input tidak valid.")
					continue
				}
				updValue := input.InputUser("Masukan data baru: ")
				fmt.Println("")
				if updValue == "q" {
					fmt.Println()
					continue
				}

				service.UpdateContact(numOpsiSelect, updOpsiSelect, updValue)
			} else {
				fmt.Println("kontak yang dipilih tidak ada")
				fmt.Println()
			}

		case "4":
			service.ReadContact()
			fmt.Println("----------------------------------------------")
			fmt.Println("q. untuk kembali ke menu utama")
			fmt.Println("----------------------------------------------")

			numStr := input.InputUser("Masukan kontak yang ingin dihapus: ")
			fmt.Println("")
			if numStr == "q" {
				fmt.Println()
				continue
			}
			numOpsiSelect, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Input tidak valid.")
				continue
			}

			service.DeleteContact(numOpsiSelect)

		case "q":
			fmt.Println()
			fmt.Println("=============================================")
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			fmt.Println("=============================================")
			return

		default:
			fmt.Println()
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			fmt.Println()
		}
	}
}
