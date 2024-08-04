package database

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/vinaocruz/go-scylla-example/config"
)

type Database struct {
	Session *gocql.Session
}

var database *Database

func NewScyllaDB(config *config.Config) *Database {
	if config == nil || config.Db == nil {
		log.Fatal("Invalid configuration")
	}

	cluster := gocql.NewCluster(config.Db.Host)
	cluster.Keyspace = config.Db.Keyspace

	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: config.Db.User,
	// 	Password: config.Db.Password,
	// }

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	database = &Database{
		Session: session,
	}

	return database
}
