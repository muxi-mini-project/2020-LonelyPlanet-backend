package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

const dns = "root:huanglingyun0130@tcp(localhost:3306)/mini_project"

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(47.97.74.180:3306)/mini_project",
		viper.GetString("db.username"),
		viper.GetString("db.password"))
	//dns := fmt.Sprintf("%s:%s@tcp(localhost:3306)/mini_project", os.Getenv("DBUser"), os.Getenv("DBPassword"))
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
		db = nil
	}
	db.SingularTable(true)
	db.LogMode(true)
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
