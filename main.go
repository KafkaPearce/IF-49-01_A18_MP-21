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

// Fungsi header untuk menampilkan header pada aplikasi
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
	fmt.Println("║                    ✦ ★ ✦ ★ ✦  🏥 SehatinPC 🏥  ✦ ★ ✦ ★ ✦                       ║")
	fmt.Println("║                                                                                ║")
	fmt.Printf("║%s%-29s%s%-34s%s║\n", lightGreen, bold, "  Go Healthier, Go Faster", reset, green)
	fmt.Println("║                                                                                ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                     ⚕️  Aplikasi Kesehatan Komputer Anda  ⚕️                     ║")
	fmt.Println("║                🔧 Maintenance • 💻 Optimization • 🛡️  Security                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Printf("%s", reset)
	fmt.Println()
}

// Fungsi menu untuk menampilkan menu utama dan menangani input pengguna
func menu(selectOption *int, cnt *int) {
	cls()
	header()
	if (*selectOption < 1 || *cnt > 2) && *cnt != 0 {
		fmt.Println("Pilihan tidak valid. Silakan pilih menu yang sesuai.")
	} else {
		fmt.Println()
	}
	fmt.Printf("%-10s[1] 🔧 Manage Component\n", "")
	fmt.Printf("%-10s[2] ↩️ Exit app\n", "")
	fmt.Println()
	fmt.Printf("%-10s╰┈➤ ", "")
	fmt.Scan(selectOption)
	*cnt++
}

// Fungsi showTable untuk menampilkan data komponen dalam bentuk tabel
func showTable(data *arrKomponen, nData int) {
	cls()
	header()
	fmt.Printf("%-5s│ %-20s│ %-15s│ %-15s│ %-15s│ %-10s│ %-10s\n", "No", "Nama Komponen", "Jenis Komponen", "Nomor Seri", "Suhu Sensor (°C)", "Beban Kerja (%)", "Status")
	fmt.Println("─────┼──────────────────────┼─────────────────┼─────────────────┼─────────────────────┼───────────────┼──────────────")
	for i := 0; i < nData; i++ {
		fmt.Printf("%-5d│ %-20s│ %-15s│ %-15s│ %-15.2f│ %-10.2f│ %-10s\n", i+1, data[i].nama, data[i].jenis, data[i].nomorSeri, data[i].suhuSensor, data[i].bebanKerja, data[i].status)
	}
}

// Fungsi createComponent untuk menambah data komponen
func createComponent(data *arrKomponen, nData *int) {
	cls()
	fmt.Println("Fitur add Belum tersedia.")
}

// Fungsi updateComponent untuk mengubah data komponen
func updateComponent(data *arrKomponen, nData *int) {
	cls()
	fmt.Println("Fitur update Belum tersedia.")
}

// Fungsi deleteComponent untuk menghapus data komponen
func deleteComponent(data *arrKomponen, nData *int) {
	cls()
	fmt.Println("Fitur delete Belum tersedia.")
}

// Fungsi manageHome untuk menampilkan menu manajemen komponen
func manageHome(data *arrKomponen, selectOption *int, cnt *int) {
	cls()
	header()
	fmt.Printf("%-10s[1] ➤ Add Component\n", "")
	fmt.Printf("%-10s[2] ➤ View Component\n", "")
	fmt.Printf("%-10s[3] ➤ Update Component\n", "")
	fmt.Printf("%-10s[4] ➤ Delete Component\n", "")
	fmt.Printf("%-10s[5] ↩️ Back to Main Menu\n", "")
	fmt.Println()
	fmt.Printf("%-10s╰┈➤ ", "")
	fmt.Scan(selectOption)
	*cnt++
}

// Fungsi main
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
			for selectOption != 5 {
				manageHome(&data, &selectOption, &cnt)
				if nData > 0 {
					switch selectOption {
					case 1:
						createComponent(&data, &nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 2:
						showTable(&data, nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 3:
						updateComponent(&data, &nData)
						fmt.Println("\nTekan Enter untuk kembali ke menu manajemen komponen...")
						fmt.Scanln()
						fmt.Scanln()
					case 4:
						deleteComponent(&data, &nData)
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
