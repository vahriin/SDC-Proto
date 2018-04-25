package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vahriin/SDC/model"
)

func ReadQuery(filename string) (string, error) {
	query, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	sQuery := string(query)
	sQuery = strings.Trim(sQuery, "\n")
	return sQuery, nil
}

func newCache(directory string) func() ([]os.FileInfo, error) {
	fileCache := make(map[string]struct{})

	return func() ([]os.FileInfo, error) {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			return nil, err
		}

		newFiles := make([]os.FileInfo, 0)

		for _, file := range files {
			if file.IsDir() {
				fileCache[file.Name()] = struct{}{}
				fmt.Println("I'm dir")
			}
			if _, ok := fileCache[file.Name()]; !ok {
				newFiles = append(newFiles, file)
				fileCache[file.Name()] = struct{}{}
			}
		}
		return newFiles, nil
	}
}

func GetQueries(directory string, qCh chan<- model.Query) {
	checkDir := newCache(directory)

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		_ = <-ticker.C
		files, err := checkDir()
		if err != nil {
			panic(err)
		}

		if len(files) > 0 {
			for _, file := range files {
				query, err := ReadQuery(directory + "/" + file.Name())

				if err == nil {
					var q model.Query
					q.Q = query
					q.T = file.ModTime()
					q.U = file.Name()

					qCh <- q
				}
			}
		}
	}
}

func GetAns(ansFile string) int64 {
	answerB, err := ioutil.ReadFile(ansFile)
	if err != nil {
		panic(err)
	}

	answer := strings.Trim(string(answerB), "\n")

	ansInt, err := strconv.ParseInt(answer, 10, 64)
	if err != nil {
		panic(err)
	}

	return ansInt
}
