package main

import (
	"github.com/vahriin/SDC/model"
)

func CheckQuery(conf *model.ConnectionConf, ansFile string, qCh <-chan model.Query, rCh chan<- model.Result) {
	processor := model.NewQueryProcessor(conf)

	defer processor.Close()

	answer := ReadAns(ansFile)

	for query := range qCh {
		var result model.Result
		result.Time = query.T
		result.Name = query.U
		if query.Q == "" {
			result.True = false
			continue
		}

		potentRes := processor.Process(query.Q)

		result.True = (potentRes == answer)

		rCh <- result
	}
}
