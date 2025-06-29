package main

import (
	"log"
	"project-1/src/config"
	"project-1/src/middlewares"
	"project-1/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Gagal memuat file .env. Menggunakan default/env sistem.")
	}

	config.ConnectDB()

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())
	router.Static("/uploads", "./uploads")

	routes.ViewRoutes(router)
	routes.AuthRoutes(router)
	routes.ContactRoutes(router)
	routes.ProductRoutes(router)
	routes.ProductAdminRoutes(router)
	routes.OrderRoutes(router)

	log.Println("🚀 Server berjalan di http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
