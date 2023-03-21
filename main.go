package main

// https://github.com/kecci/goscription/blob/master/app/cmd/http.go

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"fast-api.io/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load(filepath.Join(os.Getenv("APP_PATH"), ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fx.New(
		fx.Options(
			fx.Provide(NewGin),
			controller.Module,
		),
	).Run()
}

// func NewRecordHandler(lc fx.Lifecycle, db *gorm.DB) *handlers.RecordHandler {

// }

func NewGin(lc fx.Lifecycle) *gin.RouterGroup {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("views/**/*")
	r.Static("/assets", "./assets")

	apiRoute := r.Group("")

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Print("Starting Gin server.")
			// apiRoutes := r.Group("api")
			// apiRoutes.GET("records", handlers.RecordHandler)
			go r.Run(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Print("Stopping Gin server")
			return nil
		},
	})
	return apiRoute
}
