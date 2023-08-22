package applications

import (
	"context"
	"fmt"

	"github.com/gherbust/lab/internal/directory/domain"
	domainMongo "github.com/gherbust/lab/internal/platform/mongo/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type DirectoryMongo struct {
	DB *mongo.Client
}

func NewDirectoryMongo(db *mongo.Client) DirectoryRepository {
	return &DirectoryMongo{
		DB: db,
	}
}

func MarshallContact(contact domain.Contact) domainMongo.Contact {
	return domainMongo.Contact{
		Name:        contact.Name,
		PhoneNumber: contact.PhoneNumber,
		EMail:       contact.EMail,
		Enabled:     contact.Status(),
	}
}

func (d *DirectoryMongo) SaveContact(contact domain.Contact) {
	contact.Activation()
	mongoContacto := MarshallContact(contact)
	result, err := d.DB.Database("directorio").Collection("contacto").InsertOne(context.Background(), mongoContacto, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result, result.InsertedID)
}

func (d *DirectoryMongo) GetContact(name string) *domain.Contact {
	return nil
}

func (d *DirectoryMongo) GetAll() *[]domain.Contact {
	return nil
}

func (d *DirectoryMongo) GetAllEnabled() *[]domain.Contact {
	return nil
}

func (d *DirectoryMongo) DeleteContact(name string) {
}
