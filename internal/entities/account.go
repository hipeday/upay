package entities

type Account struct {
	Entity
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Status   string `db:"status"`
	Secret   string `db:"secret"`
}
