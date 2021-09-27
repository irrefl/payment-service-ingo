package customers

import (
	"errors"
	"github.com/google/uuid"
)

var (
	CustomerNotFoundErr = errors.New("Customer wasnt found")

	FailedWhenAddingCustomerErr = errors.New("Customer wasnt found")

	UpdateCustomerErr = errors.New("Customer wasnt found")
)

type ICustomerRepository interface {
	Get(uuid uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
