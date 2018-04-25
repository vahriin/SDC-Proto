package main

import (
	"os"
	"sort"

	"github.com/vahriin/SDC/model"
)

func WriteTop(filename string, rCh <-chan model.Result) {
	topFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer topFile.Close()

	var top model.Top

	for result := range rCh {
		top = append(top, result)
		sort.Sort(top)
		top.Write(topFile)
	}
}
