package main

import (
	"github.com/gin-gonic/gin"
	"ginlearn/controllers"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":9000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// test route gin
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// users route
	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)
	// end users route
	// banks route
	bankRepo := controllers.Baru()
	r.POST("/banks", bankRepo.CreateBank)
	r.GET("/banks", bankRepo.GetBanks)
	r.GET("/banks/:id", bankRepo.GetBankById)
	r.PUT("/banks/:id", bankRepo.UpdateBank)
	r.DELETE("/banks/:id", bankRepo.DeleteBank)
	// end banks route

	// mengembalikan nilai route
	return r
}
