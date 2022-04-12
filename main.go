package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const (
	cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"
)

func main() {
	fmt.Println("start")
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid:%d\n", syscall.Getpid())

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 200m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

}
