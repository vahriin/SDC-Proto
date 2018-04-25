package model

import (
	"fmt"
	"io"
)

type Top []Result

func (t Top) Write(w io.Writer) {
	for i, res := range t {
		fmt.Fprintf(w, "%d,%s", i+1, res.ToString())
	}
}

// implementation of Interface
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
