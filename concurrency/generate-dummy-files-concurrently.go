package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	TotalFile       = 3000
	ContentLength   = 5000
	timeoutDuration = 20 * time.Second
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "worker-pool")

type FilesInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(timeoutDuration, cancel)
	generateFilesWithContext(ctx)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	done := make(chan int)

	go func() {
		// pipeline 1: job distribution
		chanFileIndex := generateFileIndexes(ctx)

		// pipeline 2: the main logic (creating files)
		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		// track and print output
		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}

		// notify that the process is complete
		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, TotalFile)
	}
}
func generateFileIndexes(ctx context.Context) <-chan FilesInfo {
	chanOut := make(chan FilesInfo)

	go func() {
		for i := 0; i < TotalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FilesInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func randomString2(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func createFiles(ctx context.Context, chanIn <-chan FilesInfo, numberOfWorkers int) <-chan FilesInfo {
	chanOut := make(chan FilesInfo)

	//wait group to control the workers
	wg := new(sync.WaitGroup)

	//allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		//dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				//listen to chanIn channel for incoming jobs
				for job := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						//do the jobs
						filePath := filepath.Join(tempPath, job.FileName)
						content := randomString2(ContentLength)
						err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

						log.Println("workers", workerIndex, "working on", job.FileName, "file generation")

						//construct the job's result, and send it to chanOut
						chanOut <- FilesInfo{
							FileName:    job.FileName,
							WorkerIndex: workerIndex,
							Err:         err,
						}
					}
				}

				//if `chanIn` is closed, and the remaining jobs are finished,
				//only then we mark the worker as complete
				wg.Done()
			}(workerIndex)
		}
	}()

	//wait until `chanIn` closed and then all workers are done.
	//because right after that - we need to close the `chanOut` channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
