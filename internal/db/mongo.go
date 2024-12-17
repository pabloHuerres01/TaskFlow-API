package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB estructura para manejar la conexión y operaciones
type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

// ConnectMongoDB realiza la conexión a MongoDB
func ConnectMongoDB(uri, dbName, collectionName string) *MongoDB {
	// Opciones del cliente
	clientOptions := options.Client().ApplyURI(uri)

	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar a MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar conexión
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	log.Println("¡Conectado a MongoDB!")

	// Configurar base de datos y colección
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &MongoDB{
		Client:     client,
		Database:   db,
		Collection: collection,
	}
}

// InsertDocument inserta un documento en la colección
func (m *MongoDB) InsertDocument(document interface{}) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.Collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}

	log.Printf("Documento insertado con ID: %v\n", result.InsertedID)
	return result
}

// FindDocuments busca documentos que coincidan con un filtro
func (m *MongoDB) FindDocuments(filter bson.M) []bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Realizar la consulta
	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Error al buscar documentos: %v", err)
	}
	defer cursor.Close(ctx)

	// Leer resultados
	var results []bson.M
	for cursor.Next(ctx) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			log.Fatalf("Error al decodificar documento: %v", err)
		}
		results = append(results, document)
	}

	return results
}

// Close cierra la conexión con MongoDB
func (m *MongoDB) Close() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error al cerrar la conexión con MongoDB: %v", err)
	}
	log.Println("Conexión con MongoDB cerrada")
}
