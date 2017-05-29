package hn

import (
	"github.com/boltdb/bolt"
	"encoding/json"
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
		AssertResult(putResult)
		result = putResult == nil
	}
	return
}

func (this *TDataOp) ReadUser(name string) (result TUser) {
	var data = this.tx.Bucket(DataKeyUsers).Get([]byte(name))
	if (data != nil) {
		var decodeResult = json.Unmarshal(data, &result)
		AssertResult(decodeResult)
	}
	return
}

func (this *TDataOp) Login(user TUser) (SessionKey string) {
	var serverUser = this.ReadUser(user.name)
	if serverUser.CheckValid() {
		if user.password == serverUser.password {
			serverUser.SessionKey = serverUser.NewSessionKey()
			this.WriteUser(serverUser)
			SessionKey = serverUser.SessionKey
		}
	}
	return
}

func (this *TDataOp) WriteUser(user TUser) {
	this.tx.Bucket(DataKeyUsers).Put(user.GetNameBytes(), user.GetData())
}