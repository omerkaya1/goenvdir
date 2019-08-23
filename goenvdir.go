package main

import "github.com/omerkaya1/goenvdir/cmd"

//go:generate mockgen -source=internal/exec.go -destination=internal/exec_mock.go -package=internal

func main() {
	cmd.Execute()
}
