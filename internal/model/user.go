package models

type Usuario struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre string `json:"nombre" bson:"nombre"`
	Email  string `json:"email" bson:"email"`
}
