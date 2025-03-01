package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Version of the service
const Version = "3.2.1"

func main() {
	// Load cfg
	log.Printf("===> DPG Orders Service is starting up <===")
	cfg := getConfiguration()
	svc := initializeService(Version, cfg)

	// Set routes and start server
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.Default()

	// Set routes and start server
	router.Use(cors.Default())
	router.GET("/version", svc.getVersion)
	router.GET("/healthcheck", svc.healthCheck)
	router.GET("/authenticate", svc.authenticate)
	api := router.Group("/api")
	{
		api.GET("/constants", svc.getConstants)
		api.GET("/users/:id", svc.getUser)
		api.POST("/users", svc.updateUser)
		api.POST("/users/:id/address", svc.updateUserAddress)
		api.POST("/users/:id/submit", svc.submitRequest)
	}

	// Note: in dev mode, this is never actually used. The front end is served
	// by yarn and it proxies all requests to the API to the routes above
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	// add a catchall route that renders the index page.
	// based on no-history config setup info here:
	//    https://router.vuejs.org/guide/essentials/history-mode.html#example-server-configurations
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	portStr := fmt.Sprintf(":%d", cfg.port)
	log.Printf("INFO: start DPG Orders Service on port %s with CORS support enabled", portStr)
	log.Fatal(router.Run(portStr))
}
