package app

import (
	"golang-microservices/src/api/controllers/ping"
	"golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/ping", ping.Pong)
	router.POST("/repository", repositories.CreateRepo)
}
