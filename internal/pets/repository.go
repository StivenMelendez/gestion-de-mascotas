package pets

import (
	"database/sql"
	"time"
)

type Pet struct {
	ID        int       `json:"id"`
	Nombre      string    `json:"name"`
	Breed     string    `json:"breed"`
	Weight    float64   `json:"weight"`
	OwnerID   int       `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PetRepository struct {
	db *sql.DB
}

func NewPetRepository(db *sql.DB) *PetRepository {
	return &PetRepository{db: db}
}

func (r *PetRepository) Save(pet *Pet) error {
	query := `INSERT INTO Mascotas (Raza, Peso, ID_Dueno, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, pet.Breed, pet.Weight, "pet.OwnerID: " 1 /* para ejemplo */, time.Now(), time.Now())
	return err
}