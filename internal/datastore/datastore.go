package datastore

import (
	"fmt"

	"github.com/Gggdevv/go-web-template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _dataStore *DataStore

type DataStore struct {
	MySQL *gorm.DB
}

func Init(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.UserName, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	_dataStore = &DataStore{
		MySQL: db,
	}

	return nil
}

func Get() *DataStore {
	return _dataStore
}
