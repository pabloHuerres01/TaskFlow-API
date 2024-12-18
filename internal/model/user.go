package model

// Usuario representa un usuario en la base de datos.
type Usuario struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre string `json:"nombre" bson:"nombre" validate:"required"`
	Email  string `json:"email" bson:"email" validate:"required,email"`
}
