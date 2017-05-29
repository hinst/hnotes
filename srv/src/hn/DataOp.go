package hn

import (
	"github.com/boltdb/bolt"
)

type TDataOp struct {
	tx *bolt.Tx
}

func (this *TDataOp) Create(transaction *bolt.Tx) *TDataOp {
	this.tx = transaction
	return this
}

func (this *TDataOp) EnsureBuckets() {
	this.tx.CreateBucketIfNotExists(DataKeyUsers)
}

func (this *TDataOp) CheckUserExists(user TUser) bool {
	var userData = this.tx.Bucket(DataKeyUsers).Get(user.GetNameBytes())
	return userData != nil
}

func (this *TDataOp) AddNewUser(user TUser) (result bool) {
	if user.CheckValid() && false == this.CheckUserExists(user) {
		var putResult = this.tx.Bucket(DataKeyUsers).Put(user.GetNameBytes(), user.GetData())
		result = putResult == nil
	}
	return
}