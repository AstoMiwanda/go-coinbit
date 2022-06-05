package lib

import (
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	conn, err := leveldb.OpenFile("balance.db", nil)
	if err != nil {
		logrus.Errorf("Error connection db: %v", err)
	}
	//defer conn.Close()

	db = conn
}

func ConnectionDB() *leveldb.DB {
	return db
}
