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
		case "child":
			child(anotherArgs)
		default:
			println("Unknown command: ", command)
	}
}


func run(args []string) {	
	fmt.Printf("Running %v as %d\n", args, os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func child(args []string) {
	fmt.Printf("Running %v as %d\n", args, os.Getpid())
	
	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/home/ahmed/ubuntu-light")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	syscall.Unmount("/proc", 0)
}