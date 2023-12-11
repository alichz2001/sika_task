package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sika/handler"
	"sika/service"
)

var (
	app         *fiber.App
	dbConn      *gorm.DB
	userService *service.UserServiceImpl
	userHandler *handler.UserHandlerImpl
)

func init() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("missing POSTGRES_DSN env")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}
	dbConn = db
}

func init() {
	app = fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())
}

func init() {
	userService = service.NewUserService(dbConn)
	userHandler = handler.NewUserHandler(app, userService)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
