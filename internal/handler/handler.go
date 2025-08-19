package handler

import (
	"CarRentalService/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)



type Handler struct {
	service *service.Service
	redisClient *redis.Client
}

func NewHandler(service *service.Service, redisClient *redis.Client) *Handler {
	return &Handler{
		service: service,
		redisClient: redisClient,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getInfo)
			user.PUT("/", h.updateInfo)
			user.DELETE("/", h.deleteUser)
		}
		car := api.Group("/car")
		{
			car.POST("/", h.addCar)
			car.GET("/", h.getAllCars)
			car.GET("/:id", h.getCarById)
			car.DELETE("/:id", h.deleteCar)
		}
		rental := api.Group("/rental")
		{
			rental.POST("/", h.startRental)
			rental.PUT("/:id", h.endRendal)
			rental.GET("/", h.rentalHistory)
		}
		review := api.Group("/review")
		{
			review.POST("/", h.leaveReview)
		}
	}
	return router
}