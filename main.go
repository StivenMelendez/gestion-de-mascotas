package main

import (
	"fmt"
	"log"

	"gestion-de-mascotas/routes"
	"gestion-de-mascotas/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Inicializando la aplicación...")

	/*var (
		path_tipos = "default-info/tipos-mascotas.json"
		path_razas = "default-info/razas-mascotas.json"
	)*/

	e := echo.New()
	routes.RegisterRoutes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	if err := utils.EnsureUploadsDir(); err != nil {
		log.Fatalf("Error al crear el directorio de uploads: %v", err)
	}

	/*db := database.GetCollection("mascotas")

	mascotaService := mascota_service.NewMascotaService(db)*/

	//routes.RegisterRoutes(e /*, mascotaService*/)

	/*if err := tipo_service.Set(path_tipos); err != nil {
		log.Fatalf("Error al insertar los tipos de mascota: %v", err)
	}
	fmt.Println("Tipos de mascota insertados correctamente.")

	if err := raza_service.Set(path_razas); err != nil {
		log.Fatalf("Error al insertar las razas: %v", err)
	}
	fmt.Println("Razas insertadas correctamente.")*/

	port := ":7000"
	fmt.Printf("Servidor iniciado en http://localhost%s\n", port)
	log.Fatal(e.Start(port))

}
