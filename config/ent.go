package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/go-inventory-grpc/ent"

	_ "github.com/lib/pq"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

const (
	host1     = "localhost"
	port1     = 5432
	user1     = "postgres"
	password1 = "root"
	dbname1   = "inventory"
)

func NewEntClient() (*ent.Client, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host1, port1, user1, password1, dbname1)

	client, err := ent.Open("postgres", psqlInfo, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
