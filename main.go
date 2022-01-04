package main

import (
	"gin_vue/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	common.InitDB()
	db := common.GetDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	r := gin.Default()
	r = CollectRoute(r)

	panic(r.Run()) //Listen and serve on 127.0.0.1

}
