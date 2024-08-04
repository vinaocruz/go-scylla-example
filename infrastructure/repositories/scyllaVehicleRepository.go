package repositories

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/domain/entities"
	"github.com/vinaocruz/go-scylla-example/domain/repositories"
)

type ScyllaVehicleRepository struct {
	session *gocql.Session
}

func NewScyllaVehicleRepository(db *database.Database) repositories.VehicleRepository {
	return &ScyllaVehicleRepository{
		session: db.Session,
	}
}

func (r *ScyllaVehicleRepository) Load(cnh string, licensePlate string) (entities.Driver, error) {
	var driver entities.Driver
	m := map[string]interface{}{}

	iter := r.session.Query("SELECT * FROM drivers WHERE cnh = ? and license_plate = ?", cnh, licensePlate).Iter()

	if iter.NumRows() == 0 {
		return driver, errors.New("vehicle not found")
	}

	iter.MapScan(m)
	driver = entities.Driver{
		ID:           m["id"].(gocql.UUID),
		Cnh:          m["cnh"].(string),
		LicensePlate: m["license_plate"].(string),
		Name:         m["name"].(string),
		Model:        m["model"].(string),
		CreatedAt:    m["createdat"].(time.Time),
	}

	return driver, nil
}

func (r *ScyllaVehicleRepository) Delete(driver entities.Driver) error {
	return r.session.Query("DELETE FROM drivers WHERE cnh = ? and license_plate = ?", driver.Cnh, driver.LicensePlate).Exec()
}
