package interfaces

type UserModelInterface interface {
	SetName(name string)
	GetName() string
	SetRemoteAddress(ip string)
	GetRemoteAddress() string
}
