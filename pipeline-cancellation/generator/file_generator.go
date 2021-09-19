package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join("temp")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	letters := []rune("abcdefjhijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generaeFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}

		close(chanOut)
	}()

	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorker int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorker)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorker; workerIndex++ {
			go func(workerIndex int) {
				for job := range chanIn {
					filePath := filepath.Join("temp", job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
					log.Println("worker", workerIndex, "working on", job.FileName, "file generations")

					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func GenerateFiles() {
	os.RemoveAll(tempPath)
	os.Mkdir(tempPath, os.ModeParm)

	chanFileIndex := generaeFileIndexes()
	createFileWoreker := 100
}
