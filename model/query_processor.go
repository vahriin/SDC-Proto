package model

import (
	"bytes"
	"database/sql"

	_ "github.com/lib/pq"
)

type QueryProcessor struct {
	db *sql.DB
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

func (qp QueryProcessor) Close() {
	qp.db.Close()
}

func (qp QueryProcessor) Process(query string) (result int64) {
	row := qp.db.QueryRow(query)
	var res int64
	row.Scan(&res)
	return res
}
