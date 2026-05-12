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
	// Color codes untuk terminal
	green := "\033[92m"
	lightGreen := "\033[32m"
	reset := "\033[0m"
	bold := "\033[1m"

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

// Fungsi main
func main() {
	cls()
	header()
}
