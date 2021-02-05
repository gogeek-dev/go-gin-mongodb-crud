package models

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func SetupDB() *mgo.Session {
	db, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(1)
	}
	return db
}
