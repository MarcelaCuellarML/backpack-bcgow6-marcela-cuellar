package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MarcelaCuellarML/backpack-bcgow6-marcela-cuellar/GoTesting/Desafio-GoTesting/cmd/router"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run()

}
