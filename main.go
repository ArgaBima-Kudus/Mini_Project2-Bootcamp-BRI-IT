package main

import (
	"crm_service/middleware"
	"crm_service/modules/account"
	"crm_service/modules/customers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "root:12345@tcp(localhost:3306)/crm_service?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	customersHandler := customers.DefaultRequestHandler(db)
	r.POST("/customers",middleware.Auth, customersHandler.Create)
	r.GET("/customers",middleware.Auth, customersHandler.Read)
	r.GET("/customers/:id",middleware.Auth, customersHandler.ReadByID)
	r.PUT("/customers/:id",middleware.Auth, customersHandler.UpdateByID)
	r.DELETE("/customers/:id",middleware.Auth, customersHandler.DeleteByID)

	accountHandler := account.DefaultRequestHandler(db)
	r.POST("/register", accountHandler.Create)
	r.POST("/login", accountHandler.Login)
	r.GET("/account", accountHandler.ReadByUsername)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
