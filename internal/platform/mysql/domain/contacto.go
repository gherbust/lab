package domain

type Contact struct {
	Id          int    `field:"idcontacto"`
	Name        string `field:"name"`
	PhoneNumber string `field:"phone_number"`
	EMail       string `field:"e_mail"`
	Enabled     bool   `field:"enabled"`
}
