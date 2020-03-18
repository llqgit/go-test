package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"github.com/fvbock/endless"
)

func handler(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func main() {

	router := gin.Default()
	router.GET("/hello", handler)
	err := endless.ListenAndServe(":8282", router)

	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8282 stopped")

	os.Exit(0)
}
