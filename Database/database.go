package Database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var lock = &sync.Mutex{}

type (
	Database struct {
		ORM     *gorm.DB
		Options *Options
	}

	Options struct {
		Host     string
		Username string
		Password string
		DbName   string
		Port     string
		Charset  string
	}
)

var instance *Database

func GetDatabase(opts *Options) (error, *Database) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			if opts == nil || !(len(opts.Host) > 0 && len(opts.Username) > 0 && len(opts.DbName) > 0) {
				return errors.New("missing connection information"), nil
			}

			if len(opts.Port) == 0 {
				opts.Port = "3306"
			}

			if len(opts.Charset) == 0 {
				opts.Charset = "utf8mb4"
			}

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", opts.Username, opts.Password, opts.Host, opts.Port, opts.DbName, opts.Charset)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				return err, nil
			}
			instance = &Database{Options: opts, ORM: db}
		}
	}

	return nil, instance
}

// Setup - Set up an instance without returning it. If an instance already exists, it will be reset
func Setup(host string, username string, password string, dbName string) (err error) {
	if instance != nil {
		instance = nil
	}

	err, _ = GetDatabase(&Options{Host: host, Username: username, Password: password, DbName: dbName})
	return
}
