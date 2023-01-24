package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gormCrud/Migrate"
	models "github.com/gormCrud/Models"
)

func main() {
	Migrate.IntializeDb()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.POST("/product/create", Createprodcut)
	r.Run(":3030")
}

func Createprodcut(c *gin.Context) {

	db := Migrate.IntializeDb()
	var product = models.Product{}

	err := json.NewDecoder(c.Request.Body).Decode(&product)
	if err != nil {
		log.Fatal("error found in decoding data")
		panic(err)
	}
	println("code", product.Code)
	println("price", product.Price)
	db.Create(&models.Product{Code: product.Code, Price: product.Price})

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created Successfully",

		"data": product})
}
