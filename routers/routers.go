package routers

import (
	db "api/database"
	handlers "api/handlers"
	middlewares "api/middlewares"
	rates "api/rates"
	socket "api/websocket"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Router(db db.DBIInterface, dao *db.DataService) {
	router := gin.New()
	rateRepository := rates.NewRepository(dao.CubesCollection)
	rateUsecase := rates.NewUsecase(rateRepository, &http.Client{})
	ratesHandler := rates.NewHandler(rateUsecase)
	ratesHandler.SyncData()
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/rates/*any", ratesHandler.Rates)
	// router.POST("/random-rates", handlers.GetRandomRates)
	// router.GET("/value-per-currency", handlers.GetPropertyOfAll)

	router.GET("/ws", socket.Echo)

	router.POST("/mail/send-all", handlers.SendToAllUser)
	router.GET("/emails", handlers.GetAllEmails)
	router.PUT("/emails/:id", handlers.TurnOffRemindEmail)

	router.POST("/auth/login", handlers.Login)
	router.POST("/auth/register", handlers.Register)

	router.GET("/users", middlewares.CheckJwt(handlers.GetAllUsers))
	router.PUT("/users/roles/:id", middlewares.CheckJwt(handlers.SetRolesUser))
	router.PUT("/users/active/:id", middlewares.CheckJwt(handlers.ToggleUser))
	// router.Run(":" + os.Getenv("PORT"))
	router.Run(":8080")
}
