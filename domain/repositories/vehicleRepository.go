package repositories

import "github.com/vinaocruz/go-scylla-example/domain/entities"

type VehicleRepository interface {
	Load(cnh string, license_plate string) (entities.Driver, error)
	Delete(driver entities.Driver) error
}
