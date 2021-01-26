package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "contoh"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	fmt.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// dynamic
	insertDynStmt := `insert into "basket_a"("a", "fruit_a") values($1, $2)`
	_, e := db.Exec(insertDynStmt, 8, "monkey")
	CheckError(e)

	deleteStmt := `delete from "basket_a" where a=$1`
	_, e = db.Exec(deleteStmt, 1)
	CheckError(e)

}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
