package database

import (
    "budget-tracker-api/internal/config"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

func Connect() *mongo.Client {
    client, err := mongo.NewClient(options.Client().ApplyURI(GenerateURILink()))
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

func GenerateURILink() string {
	return fmt.Sprintf(
        "mongodb://%s:%s@%s:%s/%s",
        config.Get("DATABASE_USER"),
        config.Get("DATABASE_PASSWORD"),
        config.Get("DATABASE_HOST"),
        config.Get("DATABASE_PORT"),
        config.Get("DATABASE_NAME"),
    )
}