package services

import (
	"context"
	"log"
	"taskflow-api/internal/db"
	models "taskflow-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *UsuarioService) ObtenerUsuarios(ctx context.Context) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usuario models.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			log.Println("Error al decodificar usuario:", err)
			continue
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (s *UsuarioService) CrearUsuario(ctx context.Context, usuario models.Usuario) error {
	_, err := s.collection.InsertOne(ctx, usuario)
	if err != nil {
		return err
	}
	return nil
}
