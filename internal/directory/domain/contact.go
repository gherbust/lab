package domain

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	EMail       string `json:"e_mail"`
	isActive    bool
}

func (c *Contact) Activation() {
	c.isActive = true
}

func (c *Contact) Deactivation() {
	c.isActive = false
}

func (c *Contact) Status() bool {
	return c.isActive
}
