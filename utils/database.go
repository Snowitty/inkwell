package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(configPath string) error {

	viper.SetConfigFile(configPath)

	errConfig := viper.ReadInConfig()
	if errConfig != nil {
		return errConfig
	}

	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.username")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if errDB != nil {
		return errDB
	}

	return nil

}
