package cassandra

import (
	"github.com/gocql/gocql"
	"os"
)

const (
	cassandraOAuthUsername = "cassandra_oauth_username"
	cassandraOAuthPassword = "cassandra_oauth_password"
	cassandraOAuthHost     = "cassandra_oauth_host"
	cassandraOAuthKeyspace = "cassandra_oauth_keyspace"
)

var (
	username = os.Getenv(cassandraOAuthUsername)
	password = os.Getenv(cassandraOAuthPassword)
	host     = os.Getenv(cassandraOAuthHost)
	keyspace = os.Getenv(cassandraOAuthKeyspace)
	cluster  *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	defer session.Close()
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
