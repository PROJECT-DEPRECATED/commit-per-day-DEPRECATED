package main

import (
	"os"
	"os/exec"
	"time"
)

func runCode(name string, args ...string) {
	exec.Command(name, args...).Run()
}

func editFile() {
	str := "- "
	today := time.Now().Format("1/2/2006")

	str += today
	fs, err := os.ReadFile("README.md")
	if err != nil {
		return
	}

	converted := string(fs[:])
	converted += str

	convert := []byte(converted)

	os.WriteFile("README.md", convert, os.FileMode(775))
}

func push() {
	runCode("git", "add", ".")
	runCode("git", "commit", "-m", "\"Update README.md\"")
	runCode("git", "push", "origin", "master")
}

func main() {
	for {
		if time.Now().Local().Hour() == 0 {
			editFile()
			push()
		}
	}
}
