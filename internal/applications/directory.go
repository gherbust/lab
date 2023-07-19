package applications

import "github.com/gherbust-meli/lab/internal/domain"

type Directory struct {
	Storage map[string]interface{}
}

func NewDirectory() *Directory {
	storage := make(map[string]interface{}, 0)
	return &Directory{
		Storage: storage,
	}
}

func (d *Directory) SaveContact(contact domain.Contact) {
	contact.Activation()
	d.Storage[contact.Name] = contact
}

func (d *Directory) GetContact(name string) *domain.Contact {
	contact := d.Storage[name]
	if contact == nil {
		return nil
	}

	c := contact.(domain.Contact)
	return &c
}
