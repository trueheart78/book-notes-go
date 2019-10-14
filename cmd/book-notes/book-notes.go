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
	if isCreate() || isImport() {
		if hasFilename() {
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
	fmt.Println("Hello! Let's " + mode() + " " + passedFilename() + "!")
	fmt.Println("adjective: " + adjective.Get())
}

func mode() string {
	if isCreate() {
		return "create"
	} else if isImport() {
		return "import"
	}

	// return an error?
	return ""
}

func isCreate() bool {
	return len(os.Args) > 1 && os.Args[1] == "create"
}

func isImport() bool {
	return len(os.Args) > 1 && os.Args[1] == "import"
}

func hasFilename() bool {
	return len(os.Args) > 2
}

func passedFilename() string {
	if hasFilename() {
		return os.Args[2]
	}

	// return an error if hasFilename is false
	return ""
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
