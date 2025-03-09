package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var args = os.Args[1:]

	if len(args) == 0 {
		println("No arguments provided")
		return
	}

	var command = args[0]
	var anotherArgs = args[1:]

	switch command {
		case "run":
			run(anotherArgs)
		default:
			println("Unknown command: ", command)
	}
}


func run(args []string) {	
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}


	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}