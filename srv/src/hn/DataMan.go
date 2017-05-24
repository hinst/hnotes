package hn

import (
	"github.com/boltdb/bolt"
)

type TDataMan struct {
	db *bolt.DB
}

func (this *TDataMan) Start() {
	var db, dbResult = bolt.Open("data.db", 0600, nil)
	if dbResult == nil {
		this.db = db
	} else {
		panic(dbResult)
	}
}

func (this *TDataMan) Stop() {
	this.db.Close()
}