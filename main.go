package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const nmax int = 100

type komponen struct {
	nama, jenis string
	nomorSeri   string
	suhuSensor  float64
	bebanKerja  float64
	status      string
}
type arrKomponen [nmax]komponen

func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// header untuk menampilkan header pada aplikasi
func header() {
	var green, lightGreen, reset, bold string
	green = "\033[92m"
	lightGreen = "\033[32m"
	reset = "\033[0m"
	bold = "\033[1m"

	// Header dengan dekorasi
	fmt.Printf("%s", green)
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                                ║")
	fmt.Println("║                       ✦ ★ ✦ ★ ✦  🖥️   SehatinPC 🖥️  ✦ ★ ✦ ★ ✦                    ║")
	fmt.Println("║                                                                                ║")
	fmt.Printf("║%s%-29s%s%-34s%s║\n", lightGreen, bold, "      Go Aware, Go Safety", reset, green)
	fmt.Println("║                                        UHUY!                                   ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                     🤖  Aplikasi Kesehatan Komputer Anda  🤖                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Printf("%s", reset)
	fmt.Println()
}

// Fungsi menu
func menu(selectOption *int, cnt *int) {
	cls()
	header()
	if (*selectOption < 1 || *cnt > 2) && *cnt != 0 {
		fmt.Println("Pilihan tidak valid. Silakan pilih menu yang sesuai.")
	} else {
		fmt.Println()
	}
	fmt.Printf("%-10s[1] 🔧 Manage Component\n", "")
	fmt.Printf("%-10s[2] ↩️  Exit app\n", "")
	fmt.Println()
	fmt.Printf("%-10s╰┈➤ ", "")
	fmt.Scan(selectOption)
	*cnt++
}

// showTable untuk menampilkan data komponen dalam bentuk tabel
func showTable(data *arrKomponen, nData int) {
	fmt.Printf("%-5s│ %-20s│ %-15s│ %-15s│ %-15s│ %-10s│ %-10s\n", "No", "Nama Komponen", "Jenis Komponen", "Nomor Seri", "Suhu Sensor (°C)", "Beban Kerja (%)", "Status")
	fmt.Println("─────┼─────────────────────┼────────────────┼────────────────┼─────────────────┼────────────────┼──────────────")
	for i := 0; i < nData; i++ {
		fmt.Printf("%-5d│ %-20s│ %-15s│ %-15s│ %-16.2f│ %-15.2f│ %-10s\n", i+1, data[i].nama, data[i].jenis, data[i].nomorSeri, data[i].suhuSensor, data[i].bebanKerja, data[i].status)
	}
}

// statistik komponen
func statKom(data *arrKomponen, nData int) {
	var i, totalWarning int
	var totalSuhu float64
	for i = 0; i < nData; i++ {
		if data[i].status == "Warning" {
			totalWarning++
		}
		totalSuhu += data[i].suhuSensor
	}
	fmt.Println("Status: ")
	fmt.Printf("Jumlah Komponen yang bersmasalah: %d\n", totalWarning)
	fmt.Printf("Rata-rata suhu Komponen keseluruhan: %.2f\n", totalSuhu/float64(nData))
}

// createComponent untuk menambah data komponen
func createComponent(data *arrKomponen, nData *int) {
	cls()
	var newKomponen komponen
	fmt.Print("Masukkan nama komponen: ")
	fmt.Scan(&newKomponen.nama)
	fmt.Print("Masukkan jenis komponen: ")
	fmt.Scan(&newKomponen.jenis)
	fmt.Print("Masukkan nomor seri: ")
	fmt.Scan(&newKomponen.nomorSeri)
	fmt.Print("Masukkan suhu sensor: ")
	fmt.Scan(&newKomponen.suhuSensor)
	fmt.Print("Masukkan beban kerja: ")
	fmt.Scan(&newKomponen.bebanKerja)
	if newKomponen.suhuSensor > 80.0 || newKomponen.bebanKerja > 90.0 {
		newKomponen.status = "Warning"
	} else {
		newKomponen.status = "Normal"
	}
	data[*nData] = newKomponen
	(*nData)++
}

// updateComponent untuk mengubah data komponen
func updateComponent(data *arrKomponen, nData *int) {
	cls()
	var updateOption, selectOption int
	showTable(data, *nData)
	fmt.Printf("Masukkan nomor komponen yang ingin diupdate (1-%d): ", *nData)
	fmt.Scan(&updateOption)
	if updateOption >= 1 && updateOption <= *nData {
		fmt.Printf("Pilih atribut yang ingin diupdate:\n")
		fmt.Printf("[1] Nama Komponen\n")
		fmt.Printf("[2] Jenis Komponen\n")
		fmt.Printf("[3] Nomor Seri\n")
		fmt.Printf("[4] Suhu Sensor (°C)\n")
		fmt.Printf("[5] Beban Kerja (%%)\n")
		fmt.Printf("Masukkan pilihan atribut (1-6): ")
		fmt.Scan(&selectOption)
		switch selectOption {
		case 1:
			fmt.Printf("Masukkan nama komponen baru: ")
			fmt.Scan(&data[updateOption-1].nama)
		case 2:
			fmt.Printf("Masukkan jenis komponen baru: ")
			fmt.Scan(&data[updateOption-1].jenis)
		case 3:
			fmt.Printf("Masukkan nomor seri baru: ")
			fmt.Scan(&data[updateOption-1].nomorSeri)
		case 4:
			fmt.Printf("Masukkan suhu sensor baru: ")
			fmt.Scan(&data[updateOption-1].suhuSensor)
		case 5:
			fmt.Printf("Masukkan beban kerja baru: ")
			fmt.Scan(&data[updateOption-1].bebanKerja)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Nomor komponen tidak valid.")
	}
	if data[updateOption-1].suhuSensor > 85.00 || data[updateOption-1].bebanKerja > 90.00 {
		data[updateOption-1].status = "Warning"
	}
}

// deleteComponent untuk menghapus data komponen
func deleteComponent(data *arrKomponen, nData *int) {
	cls()
	var deleteOption, i int
	showTable(data, *nData)
	fmt.Printf("Masukkan nomor komponen yang ingin dihapus (1-%d): ", *nData)
	fmt.Scan(&deleteOption)
	if deleteOption >= 1 && deleteOption <= *nData {
		for i = deleteOption - 1; i < *nData-1; i++ {
			data[i] = data[i+1]
		}
		*nData--
		fmt.Println("Komponen berhasil dihapus.")
	} else {
		fmt.Println("Nomor komponen tidak valid.")
	}
}

// search untuk mencari data komponen
func searchComponent(data *arrKomponen, nData int) {
	cls()
	var i, searchOption int
	var searchQuery string

	fmt.Printf("Opsi Pencarian:\n")
	fmt.Printf("[1] Nama\n")
	fmt.Printf("[2] Status\n")
	fmt.Printf("Pilih Berdasarkan apa: ")
	fmt.Scan(&searchOption)
	if searchOption == 1 {
		fmt.Print("Masukkan nama komponen yang ingin dicari: ")
		fmt.Scan(&searchQuery)
		fmt.Println()
		fmt.Printf("%-5s│ %-20s│ %-15s│ %-15s│ %-15s│ %-10s│ %-10s\n", "No", "Nama Komponen", "Jenis Komponen", "Nomor Seri", "Suhu Sensor (°C)", "Beban Kerja (%)", "Status")
		fmt.Println("─────┼─────────────────────┼────────────────┼────────────────┼─────────────────┼────────────────┼──────────────")
		for i = 0; i < nData; i++ {
			if data[i].nama == searchQuery {
				fmt.Printf("%-5d│ %-20s│ %-15s│ %-15s│ %-16.2f│ %-15.2f│ %-10s\n", i+1, data[i].nama, data[i].jenis, data[i].nomorSeri, data[i].suhuSensor, data[i].bebanKerja, data[i].status)
			}
		}
	} else if searchOption == 2 {
		fmt.Scan(&searchQuery)
		fmt.Println()
		fmt.Printf("%-5s│ %-20s│ %-15s│ %-15s│ %-15s│ %-10s│ %-10s\n", "No", "Nama Komponen", "Jenis Komponen", "Nomor Seri", "Suhu Sensor (°C)", "Beban Kerja (%)", "Status")
		fmt.Println("─────┼─────────────────────┼────────────────┼────────────────┼─────────────────┼────────────────┼──────────────")
		for i = 0; i < nData; i++ {
			if data[i].status == searchQuery {
				fmt.Printf("%-5d│ %-20s│ %-15s│ %-15s│ %-16.2f│ %-15.2f│ %-10s\n", i+1, data[i].nama, data[i].jenis, data[i].nomorSeri, data[i].suhuSensor, data[i].bebanKerja, data[i].status)
			}
		}
	} else {
		fmt.Printf("Input tidak valid!\n")
		fmt.Scanln()
	}
}

func sortComponent(data *arrKomponen, nData int) {
	cls()
	var sortOption, i, j, ascORdesc, pass int
	var temp komponen
	fmt.Println("Pilih atribut untuk mengurutkan komponen:")
	fmt.Printf("[1] Nomor Seri\n")
	fmt.Printf("[2] Suhu Sensor (°C)\n")
	fmt.Printf("[3] Beban Kerja (%%)\n")
	fmt.Print("Masukkan pilihan atribut (1-3): ")
	fmt.Scan(&sortOption)

	if sortOption >= 1 && sortOption <= 3 {
		fmt.Print("Masukkan pilihan urutan (1-Ascending, 2-Descending): ")
		fmt.Scan(&ascORdesc)
		if ascORdesc == 1 || ascORdesc == 2 {
			switch sortOption {

			case 1:
				if ascORdesc == 1 {
					i = 1
					for i <= nData-1 {
						j = i
						temp = data[j]
						for j > 0 && temp.nomorSeri < data[j-1].nomorSeri {
							data[j] = data[j-1]
							j = j - 1
						}
						data[j] = temp
						i = i + 1
					}
				} else {
					 for i = 0; i < nData-1; i++ {
						pass = i
						for j = i + 1; j < nData; j++ {
							if data[j].nomorSeri > data[pass].nomorSeri {
								pass = j
							}
						}
						temp = data[i]
						data[i] = data[pass]
						data[pass] = temp
					}
				}
			case 2:
				if ascORdesc == 1 {
					i = 1
					for i <= nData-1 {
						j = i
						temp = data[j]
						for j > 0 && temp.suhuSensor < data[j-1].suhuSensor {
							data[j] = data[j-1]
							j = j - 1
						}
						data[j] = temp
						i = i + 1
					}
				} else {
					 for i = 0; i < nData-1; i++ {
						pass = i
						for j = i + 1; j < nData; j++ {
							if data[j].suhuSensor > data[pass].suhuSensor {
								pass = j
							}
						}
						temp = data[i]
						data[i] = data[pass]
						data[pass] = temp
					}
				}
			case 3:
				if ascORdesc == 1 {
					i = 1
					for i <= nData-1 {
						j = i
						temp = data[j]
						for j > 0 && temp.bebanKerja < data[j-1].bebanKerja {
							data[j] = data[j-1]
							j = j - 1
						}
						data[j] = temp
						i = i + 1
					}
				} else {
					 for i = 0; i < nData-1; i++ {
						pass = i
						for j = i + 1; j < nData; j++ {
							if data[j].bebanKerja > data[pass].bebanKerja {
								pass = j
							}
						}
						temp = data[i]
						data[i] = data[pass]
						data[pass] = temp
					}
				}
			}
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// manageHome untuk menampilkan menu manajemen komponen
func manageHome(data *arrKomponen, selectOption *int, cnt *int, nData int) {
	cls()
	header()
	showTable(data, nData)
	statKom(data, nData)
	fmt.Printf("%-10s[1] ➤ Add Component\n", "")
	fmt.Printf("%-10s[2] ➤ Update Component\n", "")
	fmt.Printf("%-10s[3] ➤ Delete Component\n", "")
	fmt.Printf("%-10s[4] ➤ Search Component\n", "")
	fmt.Printf("%-10s[5] ➤ Sort Component\n", "")
	fmt.Printf("%-10s[6] ↩️ Back to Main Menu\n", "")
	fmt.Println()
	fmt.Printf("%-10s╰┈➤ ", "")
	fmt.Scan(selectOption)
	*cnt++
}

// main
func main() {
	var data arrKomponen
	var selectOption int = 1
	var cnt int = 0
	var nData int = 0

	dataDummy(&data, &nData)
	for selectOption != 2 {
		menu(&selectOption, &cnt)
		if selectOption == 1 {
			cnt = 0
			for selectOption != 6 {
				manageHome(&data, &selectOption, &cnt, nData)
				if cnt > 0 {
					switch selectOption {
					case 1:
						createComponent(&data, &nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 2:
						updateComponent(&data, &nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 3:
						deleteComponent(&data, &nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 4:
						searchComponent(&data, nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 5:
						sortComponent(&data, nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					}
				}
			}
		}
	}
}

// data dummy untuk mengisi data awal pada aplikasi
func dataDummy(data *arrKomponen, nData *int) {

	data[0] = komponen{"CPU", "Processor", "12345", 65.5, 75.0, "Normal"}
	data[1] = komponen{"GPU", "Graphics Card", "67890", 70.0, 80.0, "Normal"}
	data[2] = komponen{"RAM", "Memory", "54321", 40.0, 60.0, "Normal"}
	data[3] = komponen{"HDD", "Storage", "98765", 35.0, 50.0, "Normal"}
	data[4] = komponen{"Motherboard", "Mainboard", "11223", 30.0, 40.0, "Normal"}
	*nData = 5
}
