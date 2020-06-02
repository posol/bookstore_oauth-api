package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		log.Fatal(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
