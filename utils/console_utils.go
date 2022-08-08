package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func read(kind string) int {
	var s string
	for {
		fmt.Printf(" -> Enter a %s: ", kind)
		fmt.Scanln(&s)
		val, err := strconv.Atoi(s)
		if err == nil {
			return val
		}
		fmt.Printf("\033[A\033[2K\r")
	}
}

func ReadVal() int {
	return read("value")
}

func ReadPos() int {
	return read("position")
}

func PrintDiv() {
	fmt.Println("------------------------------------")
}

func ChooseOption() {
	fmt.Print(" -> Choose an option: ")
}

func InvalidOption() {
	fmt.Println("[FAIL] Invalid option!")
	fmt.Scanln()
}

func Exit() {
	fmt.Println("Exiting. Press any key to continue...")
	ClearTerminal()
}
