package database

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/vinaocruz/go-scylla-example/config"
)

func NewScyllaDB(config *config.Config) gocqlx.Session {
	if config == nil || config.Db == nil {
		log.Fatal("Invalid configuration")
	}

	cluster := gocql.NewCluster(config.Db.Host)

	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: config.Db.User,
	// 	Password: config.Db.Password,
	// }

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	return session
}
