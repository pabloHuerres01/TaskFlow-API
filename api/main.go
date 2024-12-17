package main

import (
	"log"

	"taskflow-api/internal/db" // Reemplaza con tu ruta correcta

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// URI con autenticación
	uri := "mongodb://admin:adminpassword@localhost:27017/?authSource=admin"

	// Conectar a MongoDB
	mongoClient := db.ConnectMongoDB(uri, "taskflow", "tasks")
	defer mongoClient.Close()

	// Insertar un documento
	newTask := bson.M{
		"title":       "Primera tarea",
		"description": "Descripción de la primera tarea",
		"completed":   false,
		"createdAt":   "2024-12-17",
	}

	mongoClient.InsertDocument(newTask)

	// Buscar documentos
	filter := bson.M{"completed": false}
	results := mongoClient.FindDocuments(filter)

	log.Println("Documentos encontrados:")
	for _, doc := range results {
		log.Println(doc)
	}
}
