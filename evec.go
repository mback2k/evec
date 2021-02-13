package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Map from https://gobyexample.com/collection-functions
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := Map(os.Args[1:], os.ExpandEnv)
	if len(args) < 1 {
		return
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = syscall.Exec(path, args, os.Environ())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
