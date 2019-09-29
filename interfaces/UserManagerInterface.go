package interfaces

type UserManagerInterface interface {
	AddUser(name string, remoteAddress string)
	GetUserMap() *map[uint]UserModelInterface
	GetUserCount() uint
}
