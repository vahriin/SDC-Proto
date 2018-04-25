package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/vahriin/SDC/model"
)

func flags() (string, string, string) {
	confFile := flag.String("c", "config.json", "path to your config file")
	watchDir := flag.String("d", "queries", "path to directory for watching")
	ansFile := flag.String("a", "answer.txt", "path to your answer file")

	flag.Parse()
	return *confFile, *watchDir, *ansFile
}

func main() {
	confFile, watchDir, ansFile := flags()

	connConf := model.ReadConnectionConf(confFile)

	queryCh := make(chan model.Query)
	resultCh := make(chan model.Result)

	go WatchDir(watchDir, queryCh)
	go CheckQuery(&connConf, ansFile, queryCh, resultCh)
	go WriteTop("rating.csv", resultCh)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-exit
	close(queryCh)
	close(resultCh)
}
