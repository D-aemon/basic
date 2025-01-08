package db

import (
	"basic/internal/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	sqlDB *sql.DB
}

// Connect connect to mysql
func Connect(cfg config.Config) (db DB, err error) {
	db.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB.User, cfg.DB.Passwd, cfg.DB.IP, cfg.DB.Port, cfg.DB.DBName),
	}), &gorm.Config{})
	if err != nil {
		return
	}
	db.sqlDB, err = db.DB.DB()
	if err != nil {
		return
	}
	err = db.sqlDB.Ping()
	if err != nil {
		return
	}
	//err = db.AutoMigrate()
	//if err != nil {
	//	return
	//}
	return
}

func (db DB) Close() (err error) {
	return db.sqlDB.Close()
}

func (db DB) withTransaction(f func(DB) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		db.DB = tx // db is a copied value, so change db.DB have no side effect
		return f(db)
	})
}
