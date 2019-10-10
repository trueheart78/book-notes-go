package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/trueheart78/book-notes-go/internal/pkg/version"
)

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version") {
		fmt.Println(version.Full())
		os.Exit(0)
	}
	clearScreen()
}

func main() {
	fmt.Println("Hello")
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
