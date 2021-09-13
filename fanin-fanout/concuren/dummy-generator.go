package concuren

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

const totalFile = 5000
const contentLength = 50000

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
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFileIndexes() <-chan FileInfo {
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

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)

	// allocate N of Workers
	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen chan in
				for job := range chanIn {
					filepath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filepath, []byte(content), os.ModePerm)
					log.Println("worer", workerIndex, "working on", job.FileName, "file generation")

					// construct job results and send to chanout
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				// if `chanIn` is closed, and the remaining jobs are finished,
				// only then we mark the worker as complete.
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and then all workers are done,
	// because right after that - we need to close the `chanOut` channel.

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func GenerateFile() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// pipeline1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic(creating files)
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0

	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Println("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}
