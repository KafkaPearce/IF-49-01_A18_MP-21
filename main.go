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
	fmt.Printf("╔════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║                                                                                ║")
	fmt.Printf("║          ✦ ★ ✦ ★ ✦  🏥 SehatinPC 🏥  ✦ ★ ✦ ★ ✦                            ║")
	fmt.Printf("║                                                                                ║")
	fmt.Printf"║%s%s%s%-79s%s║\n", lightGreen, bold, "  Go Healthier, Go Faster", reset, green)
	fmt.Printf("║                                                                                ║")
	fmt.Printf("╠════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  ⚕️  Aplikasi Kesehatan Komputer Anda  ⚕️                                      ║")
	fmt.Printf("║  🔧 Maintenance • 💻 Optimization • 🛡️  Security                              ║")
	fmt.Printf("╚═════
	fmt.Printf("%s", reset)
	fmt.Println()
}

// Fungsi main
func main() {
	cls()
	header()
}
