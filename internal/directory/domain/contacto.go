package domain

//cambie el nombre de contac.go a contacto.go 31-Julio-2023 4:23pm

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

func (c *Contact) Status() bool { //Esto es el Enabled
	return c.isActive
}
