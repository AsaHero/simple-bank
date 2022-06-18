package repo

type SimpleBankRepository interface{
	CreateAccount() (uint32, error)
}