package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/payments", ProcessPayment).Methods("POST")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := c.Handler(router)

	log.Println("Serwer uruchomiony na porcie 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
