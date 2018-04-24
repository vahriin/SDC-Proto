package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type QueryProcessor struct {
	db *sql.DB
}

func (qp QueryProcessor) Close() {
	qp.db.Close()
}

func (qp QueryProcessor) Process(query string) (result int64) {
	row := qp.db.QueryRow(query)
	var res int64
	row.Scan(&res)
	return res
}

func CheckQuery(conf *ConnectionConf, ansFile string, qCh <-chan Query, rCh chan<- Result) {
	processor := NewQueryProcessor(conf)
	answer := GetAns(ansFile)

	for query := range qCh {
		var result Result
		result.Time = query.T
		result.Name = query.U
		if query.Q == "" {
			result.True = false
			continue
		}

		potentRes := processor.Process(query.Q)

		result.True = (potentRes == answer)

		fmt.Println(result)

		rCh <- result
	}
}

func NewQueryProcessor(cconf *ConnectionConf) QueryProcessor {
	b := bytes.NewBufferString("user=")
	b.WriteString(cconf.User)
	b.WriteString(" ")
	b.WriteString("password=")
	b.WriteString(cconf.Password)
	b.WriteString(" ")
	b.WriteString("dbname=")
	b.WriteString(cconf.DbName)
	b.WriteString(" ")
	b.WriteString("host=")
	b.WriteString(cconf.Host)
	b.WriteString(" ")
	b.WriteString("port=")
	b.WriteString(cconf.Port)
	b.WriteString(" ")
	b.WriteString("sslmode=disable")
	db, err := sql.Open("postgres", b.String())
	if err != nil {
		panic(err)
	}

	qp := QueryProcessor{db: db}
	return qp
}
