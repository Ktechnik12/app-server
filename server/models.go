package main

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

type Customer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Payment struct {
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	CVV        string `json:"cvv"`
}

type OrderItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	Customer Customer    `json:"customer"`
	Payment  Payment     `json:"payment"`
	Items    []OrderItem `json:"items"`
	Total    float64     `json:"total"`
}
