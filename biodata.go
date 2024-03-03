package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Nama         string
	Alamat       string
	Pekerjaan    string
	AlasanGolang string
}

// Data teman-teman di kelas
var dataTeman = map[int]Teman{
	1: {"Bagas", "Bekasi", "Web Developer", "Ingin belajar Golang"},
	2: {"Lintang", "Malang", "Cloud Engineer", "Ingin memperoleh Sertifikat"},
	3: {"Reni", "Surabaya", "Mobile Developer", "Ingin menambah Ilmu pemrograman"},
}

func main() {
	// Mendapatkan argument dari terminal
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	// Mengambil nomor absen dari argumen
	nomorAbsen := args[1]

	// Mengubah nomor absen menjadi integer
	var nomorAbsenInt int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomorAbsenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat")
		return
	}

	// Memeriksa apakah nomor absen valid
	teman, ok := dataTeman[nomorAbsenInt]
	if !ok {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
		return
	}

	// Menampilkan data teman
	fmt.Println("Data Teman:")
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.AlasanGolang)
}
