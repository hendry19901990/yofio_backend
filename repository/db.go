package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnection(DB_TYPE, MYSQL_CONNECT string) (*gorm.DB, error) {
	connection, err := gorm.Open(DB_TYPE, MYSQL_CONNECT)

	if err != nil {
		return nil, err
	}

	connection.SingularTable(true)
	connection.LogMode(true)
	connection.DB().SetConnMaxLifetime(0)

	return connection, nil
}
