package pets

import (
	"errors"
	"time"
)

type Pet struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Breed  string    `json:"breed"`
	Weight float64   `json:"weight"`
	OwnerID int      `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PetService struct {
	repo PetRepository
}

func NewPetService(repo PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) CreatePet(name string, breed string, weight float64, ownerID int) (*Pet, error) {
	if name == "" || breed == "" || weight <= 0 || ownerID <= 0 {
		return nil, errors.New("invalid pet data")
	}

	pet := &Pet{
		Name:   name,
		Breed:  breed,
		Weight: weight,
		OwnerID: ownerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Save(pet)
	if err != nil {
		return nil, err
	}

	return pet, nil
}