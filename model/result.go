package model

import (
	"fmt"
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
