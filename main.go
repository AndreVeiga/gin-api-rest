package main

import (
	"routes"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	routes.HandleRequest(r)

	r.Run()
}
