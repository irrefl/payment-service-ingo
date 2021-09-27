package customers

import (
	"errors"
	"github.com/google/uuid"
	"payment-service/domain/entities"
	"payment-service/domain/valuesObjects"
)

type Customer struct {
	person        *entities.Person
	products      []*entities.Item
	transacations []valuesObjects.Transaction
}

var (
	InvalidPersonError = errors.New("A customer needs to be a real person")
)

func NewCustomer(name string) (Customer, error) {
	isEmpty := name == ""
	if isEmpty {
		return Customer{}, InvalidPersonError
	}

	person := &entities.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:        person,
		products:      make([]*entities.Item, 0),
		transacations: make([]valuesObjects.Transaction, 0),
	}, nil
}
