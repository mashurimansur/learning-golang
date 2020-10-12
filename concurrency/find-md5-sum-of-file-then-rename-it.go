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

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func proceed() {
	var temPath = filepath.Join(os.Getenv("TEMP"), "pipeline-temp")
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(temPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		//read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		//sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		//renam file
		destination := filepath.Join(temPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destination)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
	log.Printf("%d/%d file renamed", counterRenamed, counterTotal)
}
