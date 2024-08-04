package repositories

import (
	"github.com/vinaocruz/go-scylla-example/domain/entities"
)

type DriverRepository interface {
	ListAll() ([]entities.Driver, error)
	Store(driver entities.Driver) error
	Delete(cnh string) error
	Load(cnh string) ([]entities.Driver, error)
}
