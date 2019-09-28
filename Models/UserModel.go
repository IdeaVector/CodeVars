package models

type UserModel struct {
	name          string
	remoteAddress string
}

func (user UserModel) getName() string {
	return user.name
}

func (user UserModel) setName(name string) {
	user.name = name
}

func (user UserModel) getRemoteAddress() string {
	return user.remoteAddress
}

func (user UserModel) setRemoteAddress(address string) {
	user.remoteAddress = address
}
