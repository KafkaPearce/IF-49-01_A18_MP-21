package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

// Fungsi main
func main() {
	var selectOption int = 1
	var cnt int = 0

	menu(&selectOption, &cnt)
}
