package repositories

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/domain/entities"
	"github.com/vinaocruz/go-scylla-example/domain/repositories"
)

type ScyllaDriverRepository struct {
	session *gocql.Session
}

func NewScyllaDriverRepository(db *database.Database) repositories.DriverRepository {

	return &ScyllaDriverRepository{
		session: db.Session,
	}
}

func (r *ScyllaDriverRepository) ListAll() ([]entities.Driver, error) {
	var drivers []entities.Driver
	m := map[string]interface{}{}

	iter := r.session.Query("SELECT * FROM drivers").Iter()

	for iter.MapScan(m) {
		drivers = append(drivers, r.bind(m))
		m = map[string]interface{}{}
	}

	return drivers, nil
}

func (r *ScyllaDriverRepository) Store(driver entities.Driver) error {
	if err := r.session.Query(
		"INSERT INTO drivers(id, cnh, license_plate, name, model, createdat) VALUES(?, ?, ?, ?, ?, ?)",
		driver.ID, driver.Cnh, driver.LicensePlate, driver.Name, driver.Model, driver.CreatedAt,
	).Exec(); err != nil {
		return err
	}

	return nil
}

func (r *ScyllaDriverRepository) Delete(cnh string) error {
	return r.session.Query("DELETE FROM drivers WHERE cnh = ?", cnh).Exec()
}

func (r *ScyllaDriverRepository) Load(cnh string) ([]entities.Driver, error) {
	var drivers []entities.Driver
	m := map[string]interface{}{}

	iter := r.session.Query("SELECT * FROM drivers WHERE cnh = ?", cnh).Iter()
	if iter.NumRows() == 0 {
		return drivers, errors.New("driver not found")
	}

	for iter.MapScan(m) {
		drivers = append(drivers, r.bind(m))
		m = map[string]interface{}{}
	}

	return drivers, nil
}

func (r *ScyllaDriverRepository) bind(m map[string]interface{}) entities.Driver {
	return entities.Driver{
		ID:           m["id"].(gocql.UUID),
		Cnh:          m["cnh"].(string),
		LicensePlate: m["license_plate"].(string),
		Name:         m["name"].(string),
		Model:        m["model"].(string),
		CreatedAt:    m["createdat"].(time.Time),
	}
}
