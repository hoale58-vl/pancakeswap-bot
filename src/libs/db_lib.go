package libs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func getDB() (*sql.DB, error) {
	if db == nil {
		if database, err := sql.Open("sqlite3", "./database.db"); err != nil {
			log.Println("OpenDatabase")
			log.Println(err)
			return nil, err
		} else {
			db = database
		}
	}
	return db, nil
}

func RecordData(token1 string, amount1 string, token2 string, amount2 string, blockNumber string) {
	database, err := getDB()
	if err != nil {
		log.Println("OpenDatabase")
		log.Println(err)
		return
	}

	stmt, err := database.Prepare("INSERT INTO events(token1, token2, amount1, amount2, blockNumber) values(?,?,?,?,?)")
	if err != nil {
		log.Println("INSERT data")
		log.Println(err)
		return
	}

	stmt.Exec(token1, token2, amount1, amount2, blockNumber)
	fmt.Println("new record")
}
