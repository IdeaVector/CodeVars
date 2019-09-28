package Interfaces

type UserInterface interface {
	setName(name string)
	getName() string
	setRemoteAddress(ip string)
	getRemoteAddress() string
}
