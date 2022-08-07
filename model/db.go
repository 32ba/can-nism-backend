package model

import (
	"context"
	"go-ranking-api/ent"
	"go-ranking-api/ent/migrate"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DBClient *ent.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file.: %v", err)
	}
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())

	DBClient, err = ent.Open("mysql", os.Getenv("DATABASE_URL"), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	// defer DBClient.Close()
	ctx := context.Background()
	err = DBClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
