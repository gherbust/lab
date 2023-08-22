package domain

type Contact struct {
	Name        string `bson:"name"`
	PhoneNumber string `bson:"phone_number"`
	EMail       string `bson:"e_mail"`
	Enabled     bool   `bson:"enabled"`
}
