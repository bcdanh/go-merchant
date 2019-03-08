package librarian

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database Database
var coinsID []int
var err error

//Database encapsulates database
type Database struct {
	db *sql.DB
}

//Begins a transaction
func (db Database) begin() (tx *sql.Tx) {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	return tx
}

func (db Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, err := db.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stmt
}

func (db Database) query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

func taskQuery(sql string, args ...interface{}) error {
	//log.Print("inside task query")
	SQL := database.prepare(sql)
	tx := database.begin()
	_, err = tx.Stmt(SQL).Exec(args...)
	if err != nil {
		log.Println("taskQuery: ", err)
		tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return err
		}
		//log.Println("Commit successful")
	}
	return err
}

func init() {
	database.db, err = sql.Open("sqlite3", "./scrolls.db3")
	if err != nil {
		log.Println(err)
	}

	err = taskQuery(stmtCreateCoinsTable)
	if err != nil {
		log.Println(err)
	}

	rows := database.query("SELECT ID from coins")
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
		coinsID = append(coinsID, id)
		err = rows.Err()
		if err != nil {
			log.Println(err)

		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

//Close function closes this database connection
func Close() {
	database.db.Close()
}
