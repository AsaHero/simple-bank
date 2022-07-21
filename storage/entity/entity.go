package entity

type Account struct {
	Id int64 `db:"id"`
	Owner string `db:"owner"`
	Balance int64 `db:"balance"`
	Currency string `db:"currency"`
	CreatedAt string `db:"created_at"`
}

type Entry struct {

}

type Transaction struct {
	
}