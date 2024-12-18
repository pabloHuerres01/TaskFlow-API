package services

import (
	"context"
	"fmt"
	"log"
	"taskflow-api/internal/db"
	"taskflow-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioService struct {
	collection *mongo.Collection
}

func NewUsuarioService() *UsuarioService {
	client := db.ConnectMongoDB()
	collection := client.Database("my_database").Collection("usuarios")
	return &UsuarioService{collection: collection}
}

// ObtenerUsuarios obtiene todos los usuarios de la base de datos
func (s *UsuarioService) ObtenerUsuarios(ctx context.Context) ([]model.Usuario, error) {
	var usuarios []model.Usuario
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usuario model.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			log.Println("Error al decodificar usuario:", err)
			continue
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// CrearUsuario inserta un nuevo usuario en la base de datos
func (s *UsuarioService) CrearUsuario(ctx context.Context, usuario model.Usuario) error {
	_, err := s.collection.InsertOne(ctx, usuario)
	if err != nil {
		return err
	}
	return nil
}

// EliminarUsuario elimina un usuario de la base de datos por su ID
func (s *UsuarioService) EliminarUsuario(ctx context.Context, id string) error {
	// Convertir el ID de string a ObjectId
	fmt.Println(id)

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Crear el filtro para la eliminación
	filter := bson.M{"_id": idObj}

	// Eliminar el usuario de la base de datos
	_, err = s.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return err
	}

	return nil
}
