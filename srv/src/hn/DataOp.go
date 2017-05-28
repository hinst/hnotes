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