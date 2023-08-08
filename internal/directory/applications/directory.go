package applications

import "github.com/gherbust/lab/internal/directory/domain"

type Directory struct {
	Storage map[string]interface{}
}

func NewDirectory() DirectoryRepository {
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

func (d *Directory) GetAll() *[]domain.Contact {
	contacts := []domain.Contact{}

	for _, v := range d.Storage {
		contact := v.(domain.Contact)
		contacts = append(contacts, contact)

		/* guion bajo _ no toma el valor
		Parseo: convertir un objeto a otro objeto con las mismas propiedades,
		crea una estructura de un nuevo tipo*/

		/* append agrega un dato a un arreglo*/

	}
	return &contacts
}

func (d *Directory) GetAllEnabled() *[]domain.Contact {
	contacts := []domain.Contact{}

	for _, v := range d.Storage {
		contact := v.(domain.Contact)
		contacts = append(contacts, contact)
	}
	return &contacts
}

func (d *Directory) DeleteContact(name string) *domain.Contact { //Delete 4(?)
	contact := d.Storage[name]
	if contact == nil {
		return nil
	}

	c := contact.(domain.Contact)
	return &c
}
