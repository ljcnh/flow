package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *gorm.DB

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQLClient() error {
	var (
		err error
		dsn string
	)

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		return fmt.Errorf("未设置mysql密码")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		password,
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = client.Callback().Create().Before("gorm:create").Register("gorm", beforeCreate); err != nil {
		return err
	}
	if err = client.Callback().Update().Before("gorm:update").Register("gorm:update", beforeUpdate); err != nil {
		return err
	}
	if strings.ToLower(viper.GetString("log_level")) == "debug" {
		client = client.Debug()
	}

	return nil
}

func GetClient() *gorm.DB {
	return client
}
