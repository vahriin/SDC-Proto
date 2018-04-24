package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

type Result struct {
	Name string
	True bool
	Time time.Time
}

func (r Result) ToString() string {
	return fmt.Sprintf("%s,%s,%s", r.Name, strconv.FormatBool(r.True), r.Time.Format(time.RFC822))
}

func WriteTop(filename string, rCh <-chan Result) {
	topFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer topFile.Close()

	var top Top

	for result := range rCh {
		top = append(top, result)
		sort.Sort(top)
		top.Write(topFile)
	}
}

type Top []Result

func (t Top) Write(w io.Writer) {
	for i, res := range t {
		fmt.Fprintf(w, "%d,%s", i+1, res.ToString())
	}
}

func (t Top) Len() int {
	return len(t)
}

func (t Top) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Top) Less(i, j int) bool {
	if t[i].True && !t[j].True {
		return true
	} else if !t[i].True && t[j].True {
		return false
	}

	return t[i].Time.Before(t[j].Time)
}
