package Database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var lock = &sync.Mutex{}

type (
	Database struct {
		Connection *gorm.DB
		host       string
		username   string
		password   string
		dbName     string
		port       string
	}
)

var instance *Database

func GetDatabase(host string, username string, password string, dbName string) (error, *Database) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", username, password, host, "3306", dbName)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				return err, nil
			}
			instance = &Database{host: host, username: username, password: password, dbName: dbName, port: "3306", Connection: db}
		}
	}

	return nil, instance
}

// Setup - Set up an instance without returning it. If an instance already exists, it will be reset
func Setup(host string, username string, password string, dbName string) (err error) {
	if instance != nil {
		instance = nil
	}

	err, _ = GetDatabase(host, username, password, dbName)
	return
}
