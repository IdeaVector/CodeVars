package Interfaces

type UserInterface interface {
	SetName(name string)
	GetName() string
	SetRemoteAddress(ip string)
	GetRemoteAddress() string
}
