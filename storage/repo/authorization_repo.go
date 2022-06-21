package repo

type Authorization interface{
	CreateAccount() (uint32, error)
}