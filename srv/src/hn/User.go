package hn

type TUser struct {
	name string
	password string
}

func (this TUser) CheckValid() bool {
	return len(this.name) > 0
}

func (this TUser) GetNameBytes() []byte {
	return []byte(this.name)
}