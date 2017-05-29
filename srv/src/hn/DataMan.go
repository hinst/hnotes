package hn

import (
	"github.com/boltdb/bolt"
)

type TDataMan struct {
	db *bolt.DB
}

func (this *TDataMan) Create() *TDataMan {
	return this
}

func (this *TDataMan) Start() {
	var db, dbResult = bolt.Open(AppDir + "/data/data.db", 0600, nil)
	if dbResult == nil {
		this.db = db
		this.EnsureBuckets()
	} else {
		panic(dbResult)
	}
}

func (this *TDataMan) Begin(canWrite bool) *bolt.Tx {
	var transaction, result = this.db.Begin(canWrite)
	AssertResult(result)
	return transaction
}

func (this *TDataMan) EnsureBuckets() {
	var tx = this.Begin(true)
	(&TDataOp{}).Create(tx).EnsureBuckets()
	defer tx.Commit()
}

func (this *TDataMan) RegisterUser(user TUser) (result bool) {
	if user.CheckValid() {
		this.db.Update(
			func(tx *bolt.Tx) error {
				result = (&TDataOp{}).Create(tx).AddNewUser(user)
				return nil
			})
	}
	return
}

func (this *TDataMan) Stop() {
	this.db.Close()
}