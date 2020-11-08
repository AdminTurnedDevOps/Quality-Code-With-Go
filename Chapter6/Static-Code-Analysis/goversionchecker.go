package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	isVersion()
}

func isVersion() *exec.Cmd {
	version := exec.Command("go", "version")
	version.Stdout = os.Stdout

	err1 := version.Run()

	if err1 != nil {
		log.Fatal(err1)
	}

	return version
}
