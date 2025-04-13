package main

import (
	"fmt"
	"log"
	"net/http"

	raza_service "gestion-de-mascotas/services/raza.service"
	tipo_service "gestion-de-mascotas/services/tipo.service"
)

func main() {
	// Inicializar la base de datos
	fmt.Println("Inicializando la aplicación...")

	var (
		path_tipos = "default-info/tipos-mascotas.json"
		path_razas = "default-info/razas-mascotas.json"
	)

	// Insertar datos iniciales desde los archivos JSON
	err := tipo_service.Set(path_tipos)
	if err != nil {
		log.Fatalf("Error al insertar los tipos de mascota: %v", err)
	}
	fmt.Println("Tipos de mascota insertados correctamente.")

	err = raza_service.Set(path_razas)
	if err != nil {
		log.Fatalf("Error al insertar las razas: %v", err)
	}
	fmt.Println("Razas insertadas correctamente.")

	// Configurar un servidor HTTP (opcional)
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "¡Bienvenido a la API de Gestión de Mascotas!")
	})*/

	// Iniciar el servidor
	port := ":8080"
	fmt.Printf("Servidor iniciado en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
