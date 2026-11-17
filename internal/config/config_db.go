package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBModel string

const (
	DBMysqlModel  = "mysql"
	DBPgsqlModel  = "pgsql"
	DBSqliteModel = "sqlite"
)

/*
数据库配置
*/
type DB struct {
	Mode     DBModel `yaml:"mode"` // 模式 mysql pgsql sqlite
	DBName   string  `yaml:"db_name"`
	Host     string  `yaml:"host"`
	Port     int     `yaml:"port"`
	User     string  `yaml:"user"`
	Password string  `yaml:"password"`
}

/*
根据配置文件不同的模式，选择对应的链接方式
*/
func (db *DB) Dsn() gorm.Dialector {
	switch db.Mode {
	case DBMysqlModel:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.DBName,
		)
		return mysql.Open(dsn)
	case DBPgsqlModel:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			db.Host,
			db.User,
			db.Password,
			db.DBName,
			db.Port,
		)
		return postgres.Open(dsn)
	case DBSqliteModel:
		return sqlite.Open(db.DBName)
	default:
		logrus.Warnf("未配置数据库链接")
		return nil
	}
}
