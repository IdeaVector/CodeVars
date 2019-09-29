package models

type UserModel struct {
	name          string
	remoteAddress string
}

func (user UserModel) GetName() string {
	return user.name
}

func (user UserModel) SetName(name string) {
	user.name = name
}

func (user UserModel) GetRemoteAddress() string {
	return user.remoteAddress
}

func (user UserModel) SetRemoteAddress(address string) {
	user.remoteAddress = address
}
