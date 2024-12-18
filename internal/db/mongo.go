package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB establece una conexión a MongoDB
func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017").
		SetAuth(options.Credential{
			AuthSource: "admin", // Base de datos para autenticación
			Username:   "admin",
			Password:   "admin123",
		})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	// Verificar conexión
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	log.Println("Conectado a MongoDB")
	return client
}
