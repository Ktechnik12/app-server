package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var products = []Product{
	{
		ID:          1,
		Name:        "Laptop",
		Description: "Wydajny laptop do pracy i gier",
		Price:       3499.99,
		Category:    "Elektronika",
	},
	{
		ID:          2,
		Name:        "Monitor 27\" 4K",
		Description: "Monitor o wysokiej rozdzielczości",
		Price:       1899.99,
		Category:    "Elektronika",
	},
	{
		ID:          3,
		Name:        "Klawiatura Mechaniczna",
		Description: "Klawiatura RGB z przełącznikami Cherry MX",
		Price:       449.99,
		Category:    "Akcesoria",
	},
	{
		ID:          4,
		Name:        "Mysz Bezprzewodowa",
		Description: "Ergonomiczna mysz z baterią 3-letnia",
		Price:       149.99,
		Category:    "Akcesoria",
	},
	{
		ID:          5,
		Name:        "Słuchawki Bluetooth",
		Description: "Bezprzewodowe słuchawki z ANC",
		Price:       599.99,
		Category:    "Audio",
	},
	{
		ID:          6,
		Name:        "Kabel HDMI 2.1",
		Description: "Kabel HDMI 4K@60Hz o długości 2m",
		Price:       79.99,
		Category:    "Akcesoria",
	},
}

var orders []Order

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Błąd parsowania JSON", http.StatusBadRequest)
		return
	}

	if order.Customer.FirstName == "" || order.Customer.LastName == "" {
		http.Error(w, "Brakuje danych osobowych", http.StatusBadRequest)
		return
	}

	if len(order.Items) == 0 {
		http.Error(w, "Koszyk jest pusty", http.StatusBadRequest)
		return
	}

	orders = append(orders, order)

	log.Printf("Nowe zamówienie od %s %s, razem: %.2f zł",
		order.Customer.FirstName, order.Customer.LastName, order.Total)

	response := map[string]interface{}{
		"success": true,
		"message": "Płatność przetworzona pomyślnie",
		"orderId": len(orders),
		"total":   order.Total,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
