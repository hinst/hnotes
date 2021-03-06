package huser

import (
	"encoding/json"
)

type TUser struct {
	name string
	PasswordHash string
	SessionKey string
}

const UserSessionKeyLength = 32

func (this *TUser) CheckValid() bool {
	return len(this.name) > 0
}

func (this *TUser) GetNameBytes() []byte {
	return []byte(this.name)
}

func (this *TUser) GetData() []byte {
	var data, result = json.Marshal(&this)
	AssertResult(result)
	return data
}

func (this *TUser) NewSessionKey() string {
	return MakeRandomString(UserSessionKeyLength)
}

func (this *TUser) SetPassword(password string) {
	this.PasswordHash = GetPasswordHashString(password)
}