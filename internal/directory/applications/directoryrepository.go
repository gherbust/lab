package applications

import "github.com/gherbust/lab/internal/directory/domain"

type DirectoryRepository interface {
	SaveContact(contact domain.Contact)
	GetContact(name string) *domain.Contact
	GetAll() *[]domain.Contact
	GetAllEnabled() *[]domain.Contact
	DeleteContact(name string) *domain.Contact //---------------DELETE
}

//delete 1
