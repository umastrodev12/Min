package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

func main() {
	command := os.Args[1]

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing file path.")
		}
	}
}

func runTarget(target string) {
	absPath, err := filepath.Abs(target)
	if err != nil {
		fmt.Printf("Error resolving path: %s\n", absPath)
	}

	cmd := exec.Command(absPath)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Execution Failed: %s\n", err)
	}
}

func help() {
	color.Cyan("min <command> [...flags] [...args]")
	color.Blue("Version: 0.0.1 BETA")
	fmt.Printf("\n")
	fmt.Println("Commands:")

	fmt.Println("   launch    C:/Users/usr/...")
}