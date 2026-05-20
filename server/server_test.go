package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	GetProducts(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestProcessPaymentInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		strings.NewReader("{invalid"),
	)

	w := httptest.NewRecorder()

	ProcessPayment(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestProcessPaymentMissingCustomer(t *testing.T) {
	body := `{
		"customer": {
			"firstName": "",
			"lastName": ""
		},
		"items": [
			{"id":1}
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		strings.NewReader(body),
	)

	w := httptest.NewRecorder()

	ProcessPayment(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestProcessPaymentEmptyCart(t *testing.T) {
	body := `{
		"customer": {
			"firstName":"Jan",
			"lastName":"Nowak"
		},
		"items":[]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		strings.NewReader(body),
	)

	w := httptest.NewRecorder()

	ProcessPayment(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestProcessPaymentSuccess(t *testing.T) {
	body := `{
		"customer":{
			"firstName":"Jan",
			"lastName":"Nowak"
		},
		"items":[
			{"id":1}
		],
		"total":100
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		strings.NewReader(body),
	)

	w := httptest.NewRecorder()

	ProcessPayment(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestProcessPaymentInvalidTotal(t *testing.T) {
	body := `{
		"customer":{
			"firstName":"Jan",
			"lastName":"Nowak"
		},
		"items":[
			{"id":1}
		],
		"total":"not a number"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		strings.NewReader(body),
	)

	w := httptest.NewRecorder()

	ProcessPayment(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestGetProductsResponseBody(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/products",
		nil,
	)

	w := httptest.NewRecorder()

	GetProducts(w, req)

	if !strings.Contains(
		w.Body.String(),
		"Laptop",
	) {
		t.Error("response missing products")
	}
}

func TestHealthCheckBody(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	w := httptest.NewRecorder()

	HealthCheck(w, req)

	expected := `{"status":"ok"}`

	if strings.TrimSpace(
		w.Body.String(),
	) != expected {
		t.Errorf(
			"expected %s got %s",
			expected,
			w.Body.String(),
		)
	}
}

func TestProductsCount(t *testing.T) {
	if len(products) != 6 {
		t.Errorf(
			"expected 6 products got %d",
			len(products),
		)
	}
}

func TestProductFields(t *testing.T) {
	for _, p := range products {

		if p.Name == "" {
			t.Error("empty name")
		}

		if p.Price <= 0 {
			t.Error("invalid price")
		}
	}
}
