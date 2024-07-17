package entities

type Account struct {
	Entity
	Username string `db:"username"`
	Password string `db:"password"`
	Status   string `db:"status"`
	Secret   string `db:"secret"`
}
