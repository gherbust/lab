package applications

import (
	"database/sql"
	"fmt"

	"github.com/gherbust/lab/internal/directory/domain"
	domainDB "github.com/gherbust/lab/platform/mysql/domain"
)

type DirectoryMYSQL struct {
	Storage *sql.DB
}

func NewDirectoryMYSQL(storage *sql.DB) DirectoryRepository {
	return &DirectoryMYSQL{
		Storage: storage,
	}
}

func (d *DirectoryMYSQL) SaveContact(contact domain.Contact) {
	contact.Activation()
	query := "INSERT INTO `contacto`(`name`,`phone_number`,`e_mail`,`enabled`,`last_update`) values(?,?,?,?,now())"

	result, err := d.Storage.Exec(query, contact.Name, contact.PhoneNumber, contact.EMail, contact.Status())
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (d *DirectoryMYSQL) GetContact(name string) *domain.Contact {
	cnt := domain.Contact{}

	query := "SELECT idcontacto,name,phone_number,e_mail,enabled FROM directorio.contacto where name = ?"

	rows, err := d.Storage.Query(query, name)
	if err != nil {
		fmt.Println(err.Error())
	}
	contactos := []domainDB.Contact{}
	for rows.Next() {
		contacto := new(domainDB.Contact)
		err = rows.Scan(&contacto.Id, &contacto.Name, &contacto.PhoneNumber, &contacto.EMail, &contacto.Enabled)
		if err != nil {
			fmt.Println(err.Error())
		}
		contactos = append(contactos, *contacto)
	}
	if len(contactos) > 0 {
		contacto := contactos[0]
		cnt.Name = contacto.Name
		cnt.EMail = contacto.EMail
		cnt.PhoneNumber = contacto.PhoneNumber
		if contacto.Enabled {
			cnt.Activation()
		} else {
			cnt.Deactivation()
		}
	}

	return &cnt
}

func (d *DirectoryMYSQL) GetAll() *[]domain.Contact {
	query := "SELECT idcontacto,name,phone_number,e_mail,enabled FROM contacto"

	rows, err := d.Storage.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	contactos := []domainDB.Contact{}
	for rows.Next() {
		contacto := new(domainDB.Contact)
		err = rows.Scan(&contacto.Id, &contacto.Name, &contacto.PhoneNumber, &contacto.EMail, &contacto.Enabled)
		if err != nil {
			fmt.Println(err.Error())
		}
		contactos = append(contactos, *contacto)
	}

	cnts := []domain.Contact{}
	if len(contactos) > 0 {
		for _, contacto := range contactos {
			cnt := domain.Contact{}

			cnt.Name = contacto.Name
			cnt.EMail = contacto.EMail
			cnt.PhoneNumber = contacto.PhoneNumber
			if contacto.Enabled {
				cnt.Activation()
			} else {
				cnt.Deactivation()
			}
			cnts = append(cnts, cnt)
		}
	}

	return &cnts
}

func (d *DirectoryMYSQL) GetAllEnabled() *[]domain.Contact {
	query := "SELECT idcontacto,name,phone_number,e_mail,enabled FROM contacto where enabled = 1"

	rows, err := d.Storage.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	contactos := []domainDB.Contact{}
	for rows.Next() {
		contacto := new(domainDB.Contact)
		err = rows.Scan(&contacto.Id, &contacto.Name, &contacto.PhoneNumber, &contacto.EMail, &contacto.Enabled)
		if err != nil {
			fmt.Println(err.Error())
		}
		contactos = append(contactos, *contacto)
	}

	cnts := []domain.Contact{}
	if len(contactos) > 0 {
		for _, contacto := range contactos {
			cnt := domain.Contact{}

			cnt.Name = contacto.Name
			cnt.EMail = contacto.EMail
			cnt.PhoneNumber = contacto.PhoneNumber
			if contacto.Enabled {
				cnt.Activation()
			} else {
				cnt.Deactivation()
			}
			cnts = append(cnts, cnt)
		}
	}

	return &cnts
}
