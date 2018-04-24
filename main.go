package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Query struct {
	Q string
	U string
	T time.Time
}

func flags() (string, string, string) {
	confFile := flag.String("c", "", "path to your config file")
	watchDir := flag.String("d", "", "path to directory for watching")
	ansFile := flag.String("a", "", "path to your answer file")

	flag.Parse()
	return *confFile, *watchDir, *ansFile
}

func main() {
	confFile, watchDir, ansFile := flags()

	connConf := ReadConnectionConf(confFile)

	queryCh := make(chan Query)
	resultCh := make(chan Result)

	go GetQueries(watchDir, queryCh)
	go CheckQuery(&connConf, ansFile, queryCh, resultCh)
	go WriteTop("rating.csv", resultCh)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-exit
	close(queryCh)
	close(resultCh)
}
