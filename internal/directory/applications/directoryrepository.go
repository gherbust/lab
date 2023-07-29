package infrastructure

import "github.com/gherbust/lab/internal/directory/domain"

type DirectoryRepository interface {
	SaveContact(contact domain.Contact)
	GetContact(name string) *domain.Contact
	GetAll() *[]domain.Contact
}
