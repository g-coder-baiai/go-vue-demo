package main

import (
	"gin_vue/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

func main() {

	InitConfig()
	common.InitDB()
	db := common.GetDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run()) //Listen and serve on 127.0.0.1

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
