// Hey! Min Run on terminal, he is NOT an App, i'm creating .exe version, enjoy the program~

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func main() {

	prefix := string("min://")
	port := string(":8080")

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Min is Starting...")
	fmt.Println("Current Version: 0.0.1 BETA")
	fmt.Println("Prefix:", prefix)
	fmt.Println("Port:", port)
	fmt.Println("Alert: Min can have some issues because, Min has started TOO EARLY, bye~")
	fmt.Println("------------------------------------------------------------------------")

	LaunchProgram()

}

func LaunchProgram() {

	fmt.Print("Paste your executable file here:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	appPath := strings.TrimSpace(input)
	appPath = strings.ReplaceAll(appPath, "\"", "")

	if appPath == "" {
		fmt.Println("Error: executable file is null.")
		return
	}

	cmd := exec.Command(appPath)

	fmt.Printf("Opening App: %s, \n", appPath)

	err := cmd.Run() 

	if err != nil {
		fmt.Printf("Cannot Open: %v\n", err)
	}
}
