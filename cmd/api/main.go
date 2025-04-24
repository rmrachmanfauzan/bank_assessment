package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/bank_assessment/internal/db"
	"github.com/rmrachmanfauzan/bank_assessment/internal/handler"
	"github.com/rmrachmanfauzan/bank_assessment/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	
	

	db.InitDB()
	db.RunMigrations()

	// host := os.Getenv("APP_HOST")
	// if host == "" {
	// 	host = "0.0.0.0" // Default to 0.0.0.0
	// }

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080" // Default to 8080
	}

	// Combine host and port to create the address
	address := fmt.Sprintf(":%s",  port)

	e := echo.New()
	log := logrus.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":   values.URI,
				"status": values.Status,
				"remote_ip": values.RemoteIP,
				"latency": values.Latency.String(),
				"content_length": values.ContentLength,
			}).Info("request")

			return nil
		},
	}))

	userRepo := repository.NewUserRepository(db.DB)
	accountRepo := repository.NewAccountRepository(db.DB)

	userHandler := handler.NewUserHandler(userRepo)
	accountHandler := handler.NewAccountHandler(accountRepo)
	

	e.GET("/health", func(c echo.Context) error {
        return c.String(http.StatusOK, "OK")
    })
	userRoutes := e.Group("/users")
	userRoutes.POST("/daftar", userHandler.RegisterUser)
	userRoutes.GET("/:id", userHandler.FindUser)

	accountRoutes := e.Group("/accounts")
	accountRoutes.POST("/tabung", accountHandler.TopupAccount)
	accountRoutes.POST("/tarik", accountHandler.WithdrawAccount)	
	accountRoutes.GET("/saldo/:no_rekening", accountHandler.GetSaldo)

	
	if err := e.Start(address); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

