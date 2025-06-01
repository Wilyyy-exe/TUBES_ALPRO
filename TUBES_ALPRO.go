package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime" // Import runtime untuk mendeteksi sistem operasi
)

// clearScreen membersihkan layar konsol.
// Fungsi ini mendeteksi sistem operasi dan menjalankan perintah yang sesuai.
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Untuk Windows, gunakan "cmd /c cls"
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Untuk Linux/macOS, gunakan "clear"
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout // Mengarahkan output perintah ke standar output
	cmd.Run()              // Menjalankan perintah
}

// NMAX adalah konstanta yang menentukan kapasitas maksimum data film favorit yang bisa disimpan.
const NMAX int = 100

// FilmFavorit adalah struktur data (struct) yang mendefinisikan format
// untuk menyimpan informasi film favorit seorang mahasiswa.
type FilmFavorit struct {
	ID               int
	NIM              int
	Nama             string
	JudulFilmFavorit string
}

// tabFilm adalah tipe array dengan ukuran tetap NMAX untuk menyimpan objek FilmFavorit.
type tabFilm [NMAX]FilmFavorit

// Variabel global untuk menyimpan data mahasiswa dan jumlah data yang saat ini ada.
var DataFilm tabFilm
var nData int = 0
var IDc int = 1000

// generateID menghasilkan ID unik untuk setiap film.
func generateID() int {
	IDc++
	return IDc
}

// tampilanMenu menampilkan menu utama aplikasi.
func tampilanMenu() {
	clearScreen() // Membersihkan layar sebelum menampilkan menu utama
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                   ╔═════════════════════════════════╗                   ║")
	fmt.Println("                      ║                   ║ APLIKASI FILM FAVORIT MAHASISWA ║                   ║")
	fmt.Println("                      ║                   ╚═════════════════════════════════╝                   ║")
	fmt.Println("                      ║                               Dibuat Oleh:                              ║")
	fmt.Println("                      ║                         Wily Franklyn Togatorop                         ║")
	fmt.Println("                      ║                             Gracent Alfanda                             ║")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ╠═════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║                     MENU UTAMA DATABASE FILM FAVORIT                    ║")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ╠═════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Input Data Film Favorit                                          ║")
	fmt.Println("                      ║     2. Tambah, Edit, dan Hapus Data Film                                ║")
	fmt.Println("                      ║     3. Tampilkan Semua Data Film atau Cari Data Film                    ║")
	fmt.Println("                      ║     4. Urutkan Data Film                                                ║")
	fmt.Println("                      ║     5. Tampilkan Statistik Data Film                                    ║")
	fmt.Println("                      ║     0. Keluar                                              ╔════════════╣")
	fmt.Println("                      ║                                                            ║ MENU UTAMA ║")
	fmt.Println("                      ╚════════════════════════════════════════════════════════════╩════════════╝")
	fmt.Println(" ")
	fmt.Print("                      Pilih 1/2/3/4/5/0? ")
}

// main adalah titik masuk utama program.
func main() {
	var pilihan int
	for {
		tampilanMenu()
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			if nData == 0 {
				InputFilm(&DataFilm, &nData)
			} else {
				fmt.Println("                      Data sudah ada. Silakan gunakan menu 'Tambah Data Film' untuk menambah.")
				fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
				fmt.Scanln() // Menunggu input Enter
				fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
			}
		case 2:
			if nData == 0 {
				fmt.Println("                      Masukan Data Film Terlebih Dahulu")
				fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
				fmt.Scanln() // Menunggu input Enter
				fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
			} else {
				tambahEditHapusData()
			}
		case 3:
			if nData == 0 {
				fmt.Println("                      Masukan Data Film Terlebih Dahulu")
				fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
				fmt.Scanln() // Menunggu input Enter
				fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
			} else {
				BacadanCari()
			}
		case 4:
			if nData == 0 {
				fmt.Println("                      Masukan Data Film Terlebih Dahulu")
				fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
				fmt.Scanln() // Menunggu input Enter
				fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
			} else {
				pengurutanData()
			}
		case 5:
			if nData == 0 {
				fmt.Println("                      Masukan Data Film Terlebih Dahulu")
				fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
				fmt.Scanln() // Menunggu input Enter
				fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
			} else {
				StatistikDataFilm()
			}
		case 0:
			fmt.Println("                       ████████╗██╗  ██╗ █████╗ ███╗   ██╗██╗  ██╗   ██╗   ██╗ ██████╗ ██╗   ██╗")
			fmt.Println("                       ╚══██╔══╝██║  ██║██╔══██╗████╗  ██║██║ ██╔╝    ██╗ ██╔╝██╔═══██╗██║   ██║")
			fmt.Println("                          ██║   ███████║███████║██╔██╗ ██║█████╔╝      ████╔╝ ██║   ██║██║   ██║")
			fmt.Println("                          ██║   ██╔══██║██╔══██║██║╚██╗██║██╔═██╗       ██╔╝  ██║   ██║██║   ██║")
			fmt.Println("                          ██║   ██║  ██║██║  ██║██║ ╚████║██║  ██╗      ██║   ╚██████╔╝╚██████╔╝")
			fmt.Println("                          ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝      ╚═╝    ╚═════╝  ╚═════╝ ")
			return
		default:
			fmt.Println("                      Pilihan tidak valid.")
			fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
			fmt.Scanln() // Menunggu input Enter
			fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
		}
	}
}

// InputFilm mengelola proses input data film baru dari pengguna.
func InputFilm(A *tabFilm, n *int) {
	clearScreen() // Membersihkan layar sebelum input data
	var tempFilm FilmFavorit
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                          TAMBAH DATA MAHASISWA                          ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║                Masukkan NIM, Nama, dan Film Favorit Anda          ╔═════╣")
	fmt.Println("                      ║                                                                   ║ ADD ║")
	fmt.Println("                      ╚═══════════════════════════════════════════════════════════════════╩═════╝")
	fmt.Println("")

	if *n >= NMAX {
		fmt.Println("                      Kapasitas Maximal tercapai =", *n)
		fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Menunggu input Enter
		fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
		return
	}

	fmt.Print("                      Masukkan NIM: ")
	fmt.Scan(&tempFilm.NIM)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	if findIndexByNIM(tempFilm.NIM) != -1 {
		fmt.Println("                      NIM sudah ada. Masukkan NIM yang berbeda.")
		fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Menunggu input Enter
		return
	}

	fmt.Print("                      Masukkan Nama (tanpa spasi): ")
	fmt.Scan(&tempFilm.Nama)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	fmt.Print("                      Masukkan Judul Film Favorit (tanpa spasi): ")
	fmt.Scan(&tempFilm.JudulFilmFavorit)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	tempFilm.ID = generateID()
	A[*n] = tempFilm
	*n++
	fmt.Println("                      Data Telah Ditambahkan :)")
	fmt.Printf("                      ID Film Favorit: %d\n\n", tempFilm.ID)
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// TambahData menambahkan satu data film baru. Digunakan dalam sub-menu tambah/edit/hapus.
func TambahData(A *tabFilm, n *int) {
	clearScreen() // Membersihkan layar sebelum menambahkan data
	var tempFilm FilmFavorit
	fmt.Println("                      Silahkan masukan data film yang ingin kamu tambahkan:")
	fmt.Println()
	cetakData(A[:*n], *n)
	fmt.Println()

	if *n >= NMAX {
		fmt.Println("                      Kapasitas Maximal tercapai = ", *n)
		fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Menunggu input Enter
		fmt.Scanln() // Mengkonsumsi newline yang tersisa dari fmt.Scan
		return
	}

	fmt.Print("                      Masukkan NIM: ")
	fmt.Scan(&tempFilm.NIM)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	if findIndexByNIM(tempFilm.NIM) != -1 {
		fmt.Println("                      NIM sudah ada. Tidak bisa menambahkan data duplikat.")
		fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Menunggu input Enter
		return
	}

	fmt.Print("                      Masukkan Nama (tanpa spasi): ")
	fmt.Scan(&tempFilm.Nama)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	fmt.Print("                      Masukkan Judul Film Favorit (tanpa spasi): ")
	fmt.Scan(&tempFilm.JudulFilmFavorit)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	tempFilm.ID = generateID()
	A[*n] = tempFilm
	*n = *n + 1
	fmt.Println("\n                      Data Telah Ditambahkan :)")
	fmt.Printf("                      ID Film Favorit: %d\n\n", tempFilm.ID)
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// findIndexByNIM mencari indeks data mahasiswa berdasarkan NIM menggunakan pencarian sekuensial.
func findIndexByNIM(nim int) int {
	var i int
	for i = 0; i < nData; i++ {
		if DataFilm[i].NIM == nim {
			return i
		}
	}
	return -1
}

// findIndexByNama mencari indeks data mahasiswa berdasarkan Nama menggunakan pencarian sekuensial.
func findIndexByNama(nama string) int {
	var i int
	for i = 0; i < nData; i++ {
		if DataFilm[i].Nama == nama {
			return i
		}
	}
	return -1
}

// findIndexByJudul mencari indeks data mahasiswa berdasarkan Judul Film menggunakan pencarian sekuensial.
func findIndexByJudul(judul string) int {
	var i int
	for i = 0; i < nData; i++ {
		if DataFilm[i].JudulFilmFavorit == judul {
			return i
		}
	}
	return -1
}

// findIndexByID mencari indeks data mahasiswa berdasarkan ID menggunakan pencarian sekuensial.
func findIndexByID(id int) int {
	var i int
	for i = 0; i < nData; i++ {
		if DataFilm[i].ID == id {
			return i
		}
	}
	return -1
}

// cetakData menampilkan data film favorit dalam format terstruktur.
func cetakData(A []FilmFavorit, n int) {
	var i int
	fmt.Println("                      Data Film Favorit yang sudah kamu inputkan: ")
	if n == 0 {
		fmt.Println("                      Tidak ada data.")
	} else {
		for i = 0; i < n; i++ {
			fmt.Printf("                      %d. ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", i+1, A[i].ID, A[i].NIM, A[i].Nama, A[i].JudulFilmFavorit)
		}
	}
}

// tambahEditHapusData menampilkan sub-menu untuk menambah, mengedit, atau menghapus data.
func tambahEditHapusData() {
	clearScreen() // Membersihkan layar sebelum menampilkan menu tambah/edit/hapus
	var x int
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                   MENU TAMBAH, EDIT, & HAPUS DATA FILM                  ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Tambah Data Film                                                 ║")
	fmt.Println("                      ║     2. Hapus Data Film                                                  ║")
	fmt.Println("                      ║     3. Edit Data Film                                                   ║")
	fmt.Println("                      ║     4. Keluar                                                           ║")
	fmt.Println("                      ║                                                    ╔════════════════════╣")
	fmt.Println("                      ║                                                    ║ ADD, EDIT & DELETE ║")
	fmt.Println("                      ╚════════════════════════════════════════════════════╩════════════════════╝")
	fmt.Print("                      Pilih 1/2/3/4? ")
	fmt.Scan(&x)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa
	fmt.Println()

	if x == 1 {
		TambahData(&DataFilm, &nData)
	} else if x == 2 {
		hapusData(&DataFilm, &nData)
	} else if x == 3 {
		EditData(&DataFilm, &nData)
	} else {
		return
	}
}

// hapusData menghapus data mahasiswa dari array DataFilm berdasarkan ID.
func hapusData(A *tabFilm, n *int) {
	clearScreen() // Membersihkan layar sebelum menghapus data
	var idToDelete int
	var idx, i int
	cetakData(A[:*n], *n)
	fmt.Println()

	fmt.Print("                      Masukkan ID Film Favorit yang akan dihapus: ")
	fmt.Scan(&idToDelete)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	idx = findIndexByID(idToDelete)
	if idx == -1 {
		fmt.Println("                      Data yang kamu ingin hapus tidak ditemukan :(")
	} else {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("                      Data Berhasil Dihapus :)")
	}
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// EditData memperbarui data film favorit berdasarkan ID.
func EditData(A *tabFilm, n *int) {
	clearScreen() // Membersihkan layar sebelum mengedit data
	var idToEdit, idx, editChoice, newNIM int
	fmt.Println()
	cetakData(A[:*n], *n)
	fmt.Println()

	fmt.Print("                      ID Film Favorit Apa Yang Ingin Kamu Ubah? ")
	fmt.Scan(&idToEdit)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa

	idx = findIndexByID(idToEdit)
	if idx == -1 {
		fmt.Println("                      Data Tidak Ditemukan :(")
	} else {
		fmt.Printf("                      Data saat ini: ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", A[idx].ID, A[idx].NIM, A[idx].Nama, A[idx].JudulFilmFavorit)
		fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("                      ║                                                                         ║")
		fmt.Println("                      ║                          Apa yang ingin kamu ubah?                      ║")
		fmt.Println("                      ║                                                                         ║")
		fmt.Println("                      ║     1. NIM                                                              ║")
		fmt.Println("                      ║     2. Nama                                                             ║")
		fmt.Println("                      ║     3. Judul Film Favorit                                               ║")
		fmt.Println("                      ║                                                    ╔════════════════════╣")
		fmt.Println("                      ║                                                    ║ ADD, EDIT & DELETE ║")
		fmt.Println("                      ╚════════════════════════════════════════════════════╩════════════════════╝")
		fmt.Print("                      Pilih 1/2/3? ")
		fmt.Scan(&editChoice)
		fmt.Scanln() // Mengkonsumsi newline yang tersisa

		if editChoice == 1 {
			fmt.Print("                      Mau ubah NIM menjadi: ")
			fmt.Scan(&newNIM)
			fmt.Scanln() // Mengkonsumsi newline yang tersisa
			if findIndexByNIM(newNIM) != -1 && newNIM != A[idx].NIM {
				fmt.Println("                      NIM sudah ada. Tidak bisa mengubah ke NIM yang sudah ada.")
			} else {
				A[idx].NIM = newNIM
				fmt.Println("                      NIM berhasil diubah.")
			}
		} else if editChoice == 2 {
			fmt.Print("                      Mau ubah Nama menjadi (tanpa spasi): ")
			fmt.Scan(&A[idx].Nama)
			fmt.Scanln() // Mengkonsumsi newline yang tersisa
			fmt.Println("                      Nama berhasil diubah.")
		} else if editChoice == 3 {
			fmt.Print("                      Mau ubah Judul Film Favorit menjadi (tanpa spasi): ")
			fmt.Scan(&A[idx].JudulFilmFavorit)
			fmt.Scanln() // Mengkonsumsi newline yang tersisa
			fmt.Println("                      Judul Film Favorit berhasil diubah.")
		} else {
			fmt.Println("                      Pilihan tidak valid.")
		}
	}
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// BacadanCari menampilkan sub-menu untuk mencari atau mencetak semua data.
func BacadanCari() {
	clearScreen() // Membersihkan layar sebelum menampilkan menu cari/cetak
	var pilih int
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                       MENU CETAK & CARI DATA FILM                       ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Cari Data Film                                                   ║")
	fmt.Println("                      ║     2. Cetak semua Data Film                                            ║")
	fmt.Println("                      ║     3. Keluar                                                           ║")
	fmt.Println("                      ║                                                        ╔════════════════╣")
	fmt.Println("                      ║                                                        ║ PRINT & SEARCH ║")
	fmt.Println("                      ╚════════════════════════════════════════════════════════╩════════════════╝")
	fmt.Print("                      Pilih 1/2/3? ")
	fmt.Scan(&pilih)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa
	fmt.Println()
	if pilih == 1 {
		cariData(&DataFilm, nData)
	} else if pilih == 2 {
		cetakData(DataFilm[:nData], nData)
		fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Menunggu input Enter
	} else {
		return
	}
}

// cariData mencari dan menampilkan data berdasarkan kriteria tertentu.
func cariData(A *tabFilm, n int) {
	clearScreen() // Membersihkan layar sebelum mencari data
	var pilih, nim, id, idx, i int
	var nama, judul string
	var found bool

	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                       MENU CETAK & CARI DATA FILM                       ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Cari Berdasarkan ID                                              ║")
	fmt.Println("                      ║     2. Cari Berdasarkan NIM                                             ║")
	fmt.Println("                      ║     3. Cari Berdasarkan Nama                                            ║")
	fmt.Println("                      ║     4. Cari Berdasarkan Judul Film                                      ║")
	fmt.Println("                      ║     5. Keluar                                                           ║")
	fmt.Println("                      ║                                                        ╔════════════════╣")
	fmt.Println("                      ║                                                        ║ PRINT & SEARCH ║")
	fmt.Println("                      ╚════════════════════════════════════════════════════════╩════════════════╝")
	fmt.Print("                      Pilih 1/2/3/4/5? ")
	fmt.Scan(&pilih)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa
	fmt.Println()

	if pilih == 1 {
		fmt.Print("                      Masukan ID Yang Ingin Dicari: ")
		fmt.Scan(&id)
		fmt.Scanln() // Mengkonsumsi newline yang tersisa
		idx = findIndexByID(id)
		if idx == -1 {
			fmt.Println("                      Data yang kamu cari tidak ditemukan :(")
		} else {
			fmt.Printf("                      ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", A[idx].ID, A[idx].NIM, A[idx].Nama, A[idx].JudulFilmFavorit)
		}
	} else if pilih == 2 {
		fmt.Print("                      Masukan NIM Yang Ingin Dicari: ")
		fmt.Scan(&nim)
		fmt.Scanln() // Mengkonsumsi newline yang tersisa
		idx = findIndexByNIM(nim)
		if idx == -1 {
			fmt.Println("                      Data yang kamu cari tidak ditemukan :(")
		} else {
			fmt.Printf("                      ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", A[idx].ID, A[idx].NIM, A[idx].Nama, A[idx].JudulFilmFavorit)
		}
	} else if pilih == 3 {
		fmt.Print("                      Masukan Nama Yang Ingin Dicari: ")
		fmt.Scan(&nama)
		fmt.Scanln() // Mengkonsumsi newline yang tersisa
		found = false
		for i = 0; i < n; i++ {
			if A[i].Nama == nama {
				fmt.Printf("                      ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", A[i].ID, A[i].NIM, A[i].Nama, A[i].JudulFilmFavorit)
				found = true
			}
		}
		if !found {
			fmt.Println("                      Data yang kamu cari tidak ditemukan :(")
		}
	} else if pilih == 4 {
		fmt.Print("                      Masukan Judul Film Yang Ingin Dicari: ")
		fmt.Scan(&judul)
		fmt.Scanln() // Mengkonsumsi newline yang tersisa
		found = false
		for i = 0; i < n; i++ {
			if A[i].JudulFilmFavorit == judul {
				fmt.Printf("                      ID: %d, NIM: %d, Nama: %s, Film Favorit: %s\n", A[i].ID, A[i].NIM, A[i].Nama, A[i].JudulFilmFavorit)
				found = true
			}
		}
		if !found {
			fmt.Println("                      Data yang kamu cari tidak ditemukan :(")
		}
	} else {
		return
	}
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// sortByIDAsc mengurutkan data film berdasarkan ID secara ascending menggunakan Insertion Sort.
func sortByIDAsc(data []FilmFavorit, n int) {
	var pass, i int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = data[pass]
		for i > 0 && temp.ID < data[i-1].ID {
			data[i] = data[i-1]
			i--
		}
		data[i] = temp
	}
	fmt.Println("                      Pengurutan data Film dari ID terkecil: ")
	cetakData(data, n)
}

// sortByIDDesc mengurutkan data film berdasarkan ID secara descending menggunakan Selection Sort.
func sortByIDDesc(data []FilmFavorit, n int) {
	var pass, i, idx int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if data[i].ID > data[idx].ID {
				idx = i
			}
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
	}
	fmt.Println("                      Pengurutan data Film dari ID terbesar: ")
	cetakData(data, n)
}

// sortByNIMAsc mengurutkan data film berdasarkan NIM secara ascending menggunakan Insertion Sort.
func sortByNIMAsc(data []FilmFavorit, n int) {
	var pass, i int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = data[pass]
		for i > 0 && temp.NIM < data[i-1].NIM {
			data[i] = data[i-1]
			i--
		}
		data[i] = temp
	}
	fmt.Println("                      Pengurutan data Film dari NIM terkecil: ")
	cetakData(data, n)
}

// sortByNIMDesc mengurutkan data film berdasarkan NIM secara descending menggunakan Selection Sort.
func sortByNIMDesc(data []FilmFavorit, n int) {
	var pass, i, idx int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if data[i].NIM > data[idx].NIM {
				idx = i
			}
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
	}
	fmt.Println("                      Pengurutan data Film dari NIM terbesar: ")
	cetakData(data, n)
}

// sortByNamaAsc mengurutkan data film berdasarkan Nama secara ascending menggunakan Insertion Sort.
func sortByNamaAsc(data []FilmFavorit, n int) {
	var pass, i int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = data[pass]
		for i > 0 && temp.Nama < data[i-1].Nama {
			data[i] = data[i-1]
			i--
		}
		data[i] = temp
	}
	fmt.Println("                      Pengurutan data Film berdasarkan Nama (A-Z): ")
	cetakData(data, n)
}

// sortByNamaDesc mengurutkan data film berdasarkan Nama secara descending menggunakan Selection Sort.
func sortByNamaDesc(data []FilmFavorit, n int) {
	var pass, i, idx int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if data[i].Nama > data[idx].Nama {
				idx = i
			}
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
	}
	fmt.Println("                      Pengurutan data Film berdasarkan Nama (Z-A): ")
	cetakData(data, n)
}

// sortByJudulFilmAsc mengurutkan data film berdasarkan Judul Film secara ascending menggunakan Insertion Sort.
func sortByJudulFilmAsc(data []FilmFavorit, n int) {
	var pass, i int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		i = pass
		temp = data[pass]
		for i > 0 && temp.JudulFilmFavorit < data[i-1].JudulFilmFavorit {
			data[i] = data[i-1]
			i--
		}
		data[i] = temp
	}
	fmt.Println("                      Pengurutan data Film berdasarkan Judul (A-Z): ")
	cetakData(data, n)
}

// sortByJudulFilmDesc mengurutkan data film berdasarkan Judul Film secara descending menggunakan Selection Sort.
func sortByJudulFilmDesc(data []FilmFavorit, n int) {
	var pass, i, idx int
	var temp FilmFavorit
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if data[i].JudulFilmFavorit > data[idx].JudulFilmFavorit {
				idx = i
			}
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
	}
	fmt.Println("                      Pengurutan data Film berdasarkan Judul (Z-A): ")
	cetakData(data, n)
}

// pengurutanData menampilkan sub-menu untuk mengurutkan data film dan memanggil fungsi pengurutan yang sesuai.
func pengurutanData() {
	clearScreen() // Membersihkan layar sebelum menampilkan menu pengurutan
	var Pilih int

	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                          MENU PENGURUTAN DATA                           ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Pengurutan berdasarkan ID (Kecil ke Besar)                       ║")
	fmt.Println("                      ║     2. Pengurutan berdasarkan ID (Besar ke Kecil)                       ║")
	fmt.Println("                      ║     3. Pengurutan berdasarkan NIM (Kecil ke Besar)                      ║")
	fmt.Println("                      ║     4. Pengurutan berdasarkan NIM (Besar ke Kecil)                      ║")
	fmt.Println("                      ║     5. Pengurutan berdasarkan Nama (A-Z)                                ║")
	fmt.Println("                      ║     6. Pengurutan berdasarkan Nama (Z-A)                                ║")
	fmt.Println("                      ║     7. Pengurutan berdasarkan Judul Film (A-Z)                          ║")
	fmt.Println("                      ║     8. Pengurutan berdasarkan Judul Film (Z-A)                          ║")
	fmt.Println("                      ║     9. Keluar                                                           ║")
	fmt.Println("                      ║                                                                  ╔══════╣")
	fmt.Println("                      ║                                                                  ║ SORT ║")
	fmt.Println("                      ╚══════════════════════════════════════════════════════════════════╩══════╝")
	fmt.Print("                      Pilih 1/2/3/4/5/6/7/8/9? ")
	fmt.Scan(&Pilih)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa
	fmt.Println()

	var tempData []FilmFavorit = make([]FilmFavorit, nData)
	copy(tempData, DataFilm[:nData])

	switch Pilih {
	case 1:
		sortByIDAsc(tempData, nData)
	case 2:
		sortByIDDesc(tempData, nData)
	case 3:
		sortByNIMAsc(tempData, nData)
	case 4:
		sortByNIMDesc(tempData, nData)
	case 5:
		sortByNamaAsc(tempData, nData)
	case 6:
		sortByNamaDesc(tempData, nData)
	case 7:
		sortByJudulFilmAsc(tempData, nData)
	case 8:
		sortByJudulFilmDesc(tempData, nData)
	case 9:
		return
	default:
		fmt.Println("                      Pilihan tidak valid.")
	}
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}

// StatistikDataFilm menampilkan statistik data film.
func StatistikDataFilm() {
	clearScreen() // Membersihkan layar sebelum menampilkan statistik
	var pilih int
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                              MENU STATISTIK                             ║")
	fmt.Println("                      ╚═════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("                      ╔═════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║                       Apa yang ingin anda ketahui?                      ║")
	fmt.Println("                      ║                                                                         ║")
	fmt.Println("                      ║     1. Jumlah Total Data Film Yang Sudah Tersimpan                      ║")
	fmt.Println("                      ║     2. Keluar                                                           ║")
	fmt.Println("                      ║                                                                ╔════════╣")
	fmt.Println("                      ║                                                                ║ SEARCH ║")
	fmt.Println("                      ╚════════════════════════════════════════════════════════════════╩════════╝")
	fmt.Print("                      Pilih 1/2? ")
	fmt.Scan(&pilih)
	fmt.Scanln() // Mengkonsumsi newline yang tersisa
	fmt.Println()

	if pilih == 1 {
		fmt.Printf("                      Jumlah Total Data Film Yang Tersimpan: %d\n", nData)
	} else if pilih == 2 {
		return
	} else {
		fmt.Println("                      Pilihan tidak valid.")
	}
	fmt.Println("\n                      Tekan Enter untuk melanjutkan...")
	fmt.Scanln() // Menunggu input Enter
}
