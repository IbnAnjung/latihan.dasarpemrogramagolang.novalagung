package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var path = "test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("===> file berhasil fibuat", path)
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Mari belajac golang\n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("===> file berhasil di isi")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0664)
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}

		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil di hapus")
}

func main() {
	createFile()
	writeFile()
	var open, _ = exec.Command("code", "test.txt").Output()
	fmt.Printf("%s\n", string(open))
	readFile()
	deleteFile()
}
