package main

import (
	"context"
	"log"
	"os"

	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/pkg/database"
)

func main(){
	ctx := context.Background()


	// Inticialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error .env path is required")
		}
		return os.Args[1]
	}())



	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	log.Println(db)
}