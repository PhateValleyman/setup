package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

func setup(args []string) {
	// If no arguments provided or argument is "help", "-h", "-?", or "--help",
	// display the help message
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "-?" || args[0] == "--help" {
		printHelp()
		return
	}
	// If the argument is "version", "-v", or "--version", display the script version
	if args[0] == "version" || args[0] == "-v" || args[0] == "--version" {
		printVersion()
		return
	}
	// Check the provided argument and set the respective file path
	var filePath string
	switch args[0] {
	case "alias":
		filePath = filepath.Join(getPrefix(), "etc", "profile.d", "alias.sh")
	case "defaults":
		filePath = filepath.Join(getPrefix(), "etc", "profile.d", "defaults.sh")
	case "functions":
		filePath = filepath.Join(getPrefix(), "etc", "profile.d", "functions.sh")
	case "paths":
		filePath = filepath.Join(getPrefix(), "etc", "profile.d", "paths.sh")
	case "profile":
		filePath = filepath.Join(getPrefix(), "etc", "profile")
	default:
		fmt.Println("Invalid option:", args[0], ". Run 'setup' for help.")
		return
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	// Open the specified file for editing using the default editor (or nano if not set)
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error opening editor:", err)
		return
	}

	// Source the modified file (this part might need adaptation based on your requirements)
	fmt.Println("Sourcing", filePath)
}

func getPrefix() string {
	_, currentFile, _, _ := runtime.Caller(1)
	return filepath.Dir(filepath.Dir(currentFile))
}

func printHelp() {
	// Add your help message here
}

func printVersion() {
	color.Cyan("setup v1.1.0\nby PhateValleyman\nJonas.Ned@outlook.com")
}

func main() {
	setup(os.Args[1:])
}
