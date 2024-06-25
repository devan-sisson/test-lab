package main

import (
	// "fmt"
	// "io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func notMain() {
	targetDir := "WowWhatATestDir"
	// os.Chdir(os.Getenv("HOME"))
	toDir := os.Args[1]
	os.Chdir(toDir)
	err := os.Mkdir(targetDir, 0755)
	check(err)

	createEmptyFile := func(name string) {
		d := []byte(""+name)
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(filepath.Join(targetDir, "testfile"))
	createEmptyFile(filepath.Join(targetDir, "hello"))

}
