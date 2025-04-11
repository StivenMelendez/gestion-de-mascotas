package pets

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Pet struct {
	ID     int     `json:"id"`
	Raza   string  `json:"raza"`
	Peso   float64 `json:"peso"`
	ID_Dueno int   `json:"id_dueno"`
	ID_Tipo int     `json:"id_tipo"`
}

type Handler struct {
	Service *PetService
}

func NewHandler(service *PetService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) AddPet(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreatePet(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/pets", h.AddPet).Methods("POST")
}