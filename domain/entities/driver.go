package entities

import (
	"time"

	"github.com/gocql/gocql"
)

type Driver struct {
	ID           gocql.UUID `json:"id" db:"id"`
	Cnh          string     `json:"cnh" db:"cnh"`
	LicensePlate string     `json:"license_plate" db:"license_plate"`
	Name         string     `json:"name" db:"name"`
	Model        string     `json:"model" db:"model"`
	CreatedAt    time.Time  `json:"created_at" db:"createdat"`
}
