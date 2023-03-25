package main

import (
	"context"
	"fmt"
	"log"
	"os"

	config "example.com/go-inventory-grpc/config"
	staffDomain "example.com/go-inventory-grpc/internal/domain/staff"
	"example.com/go-inventory-grpc/internal/repository"
	staffRepo "example.com/go-inventory-grpc/internal/repository/staff"
	server "example.com/go-inventory-grpc/server"
)

func main() {

	ctx := context.Background()

	log.Printf("Inventory Service Initilizing Database Connection.......")
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("failed to load in configuration: ", err)
		os.Exit(1)
	}

	postgresURL := cfg.Database.PostgresURL
	if postgresURL == "" || postgresURL == "null" {
		postgresURL = config.GetPostgresURL()
	}

	db, err := repository.NewPostgres(ctx, postgresURL)
	if err != nil {
		fmt.Println("Failed to set up  DB")
	}

	staffRepo := staffRepo.New(db)

	staffD := staffDomain.New(staffRepo)

	serv := server.New(db, cfg, staffD)

	if err := serv.Start(); err != nil {
		fmt.Println("Failed to start service")
	}

	// go func() {
	// 	if err := serv.Start(); err != nil {
	// 		fmt.Println("Failed to start service")
	// 	}
	// }()

}
