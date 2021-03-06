package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PORT string

func init() {
	var err error
	viper.SetConfigName("app")
	viper.SetConfigType("properties")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n ", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No File")
		} else {
			fmt.Println("Find file but have err..")
		}
	}
	PORT = viper.GetString("server.port")
	url := viper.GetString("db.url")
	db := viper.GetString("db.databases")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	// 拼接url
	dsn := username + ":" + password + "@tcp(" + url + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
