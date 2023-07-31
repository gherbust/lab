package domain

type Contact struct {
	Id          int    `db:"idcontacto"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	EMail       string `db:"e_mail"`
	Enabled     bool   `db:"enabled"`
}
