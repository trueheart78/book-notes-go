package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/trueheart78/book-notes-go/internal/pkg/adjective"
	"github.com/trueheart78/book-notes-go/internal/pkg/version"
)

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version") {
		fmt.Println(version.Full())
		os.Exit(0)
	}
	if len(os.Args) > 1 && (os.Args[1] == "create" || os.Args[1] == "import") {
		if len(os.Args) > 2 {
			clearScreen()
		} else {
			fmt.Println("Error: Missing filename.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: Unrecognized command. Try 'create' or 'import'")
		os.Exit(1)
	}
}

func main() {
	if os.Args[1] == "create" {
		fmt.Println("Hello! Let's create " + os.Args[2] + "!")
	} else if os.Args[1] == "import" {
		fmt.Println("Hello! Let's import " + os.Args[2] + "!")
	}
	fmt.Println("adjective: " + adjective.Get())
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
