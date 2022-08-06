package main

import (
	"context"
	"go-ranking-api/controller"
	"go-ranking-api/ent"
	"go-ranking-api/ent/migrate"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file.: %v", err)
	} 
	client, err :=  ent.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
    err = client.Schema.Create(
        ctx, 
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true), 
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	engine := gin.Default()

	engine.GET("/healthchecker", controller.HealthChecker)

	engine.Run(":8080")
}