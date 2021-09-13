package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath = filepath.Join("temp")

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "second")

}

func proceed() {
	counterTotal := 0
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		sum := fmt.Sprintf("%x", md5.Sum(buf))

		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("Error", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}
