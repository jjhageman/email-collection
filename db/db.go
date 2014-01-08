package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/launch-rock/email"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB // global variable to share it between main and the HTTP handler

func InitDb(conn string) *gorp.DbMap {
	// connect to db using standard Go database/sql API
	db, err := sql.Open("postgres", conn)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(email.Email{}, "emails").SetKeys(true, "ID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
