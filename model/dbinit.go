package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

const dns = "root:huanglingyun0130@tcp(localhost:3306)/mini_project"

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(localhost:3306)/mini_project",
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"))
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
		db = nil
	}
	db.SingularTable(true)
	return db, err
}

func (db *Database) Init() {
	newDb, err := getDatabase()
	if err != nil {
		log.Printf("DBInit")
		fmt.Println(err)
	}
	Db = &Database{Self: newDb}
}

func (db *Database) Close() error {
	if err := Db.Self.Close(); err != nil {
		return err
	}
	return nil
}
