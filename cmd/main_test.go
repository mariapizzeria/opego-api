package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lib/pq"
)

func TestPostAndCancelOrderSuccess(t *testing.T) {
	tester := httptest.NewServer(App())
	defer tester.Close()

	// Создание заказа
	requestBody := strings.NewReader(`{
		"passenger_id": 1,
		"address_from": "ул. Братьев Кашириных",
		"address_to": "ТЦ МореМолл",
		"tariff": "comfort_plus",
		"selected_services": ["pet"], 
		"comment": "Буду на парковке"
	}`)

	res, err := http.Post(tester.URL+"/api/order", "application/json", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Fatalf("Expected status 201, got %d", res.StatusCode)
	}

	var response struct {
		OrderId          uint           `json:"order_id"`
		UpdatedAt        time.Time      `json:"updated_at"`
		CreatedAt        time.Time      `json:"created_at"`
		CanceledAt       *time.Time     `json:"canceled_at"`
		CompletedAt      *time.Time     `json:"completed_at"`
		PassengerId      uint           `json:"passenger,omitempty"`
		OrderStatus      string         `json:"order_status"`
		DriverAssigned   *uint          `json:"driver,omitempty"`
		ArrivedCode      string         `json:"arrived_code"`
		AddressFrom      string         `json:"address_from"`
		AddressTo        string         `json:"address_to"`
		Tariff           string         `json:"tariff"`
		SelectedServices pq.StringArray `json:"selected_services"`
		Comment          string         `json:"comment"`
		Price            int            `json:"price"`
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response.OrderId == 0 {
		t.Fatal("Order ID is empty")
	}

	t.Logf("Created order with ID: %d", response.OrderId)

	orderIDStr := strconv.FormatUint(uint64(response.OrderId), 10)

	// Отмена заказа
	res, err = http.Post(tester.URL+"/api/order/"+orderIDStr+"/cancel", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		t.Fatalf("Expected status 204, got %d", res.StatusCode)
	}

	t.Logf("Orders with ID: %d is canceled", response.OrderId)
}

func TestOrderFullLifecycleSuccess(t *testing.T) {
	tester := httptest.NewServer(App())
	defer tester.Close()

	// Создание заказа
	requestBody := strings.NewReader(`{
		"passenger_id": 1,
		"address_from": "ул. Ленина 1",
		"address_to": "ул. Пушкина 10",
		"tariff": "comfort",
		"selected_services": [], 
		"comment": ""
	}`)

	res, err := http.Post(tester.URL+"/api/order", "application/json", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Fatalf("Expected status 201, got %d", res.StatusCode)
	}

	var response struct {
		OrderId          uint           `json:"order_id"`
		UpdatedAt        time.Time      `json:"updated_at"`
		CreatedAt        time.Time      `json:"created_at"`
		CanceledAt       *time.Time     `json:"canceled_at"`
		CompletedAt      *time.Time     `json:"completed_at"`
		PassengerId      uint           `json:"passenger,omitempty"`
		OrderStatus      string         `json:"order_status"`
		DriverAssigned   *uint          `json:"driver,omitempty"`
		ArrivedCode      string         `json:"arrived_code"`
		AddressFrom      string         `json:"address_from"`
		AddressTo        string         `json:"address_to"`
		Tariff           string         `json:"tariff"`
		SelectedServices pq.StringArray `json:"selected_services"`
		Comment          string         `json:"comment"`
		Price            int            `json:"price"`
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response.OrderId == 0 {
		t.Fatal("Order ID is empty")
	}

	t.Logf("Created order with ID: %d", response.OrderId)

	orderIDStr := strconv.FormatUint(uint64(response.OrderId), 10)

	// Перевод заказа в статус search
	resp, err := http.NewRequest("PUT", tester.URL+"/api/order/"+orderIDStr+"/status/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	r, err := client.Do(resp)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()

	if r.StatusCode != 201 {
		t.Fatalf("Expected status 201, got %d", r.StatusCode)
	}

	t.Logf("Order status: search")

	// Водитель принял заказ
	res, err = http.Post(tester.URL+"/api/order/"+orderIDStr+"/accept", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected status 200 for accept, got %d", res.StatusCode)
	}

	t.Logf("Driver accepted order")

	// Водитель подъехал
	res, err = http.Post(tester.URL+"/api/order/"+orderIDStr+"/arrived", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected status 200 for arrived, got %d", res.StatusCode)
	}
	t.Logf("Order status: waiting for confirmation")

	// Водитель изменил статус заказа на in_progress
	requestBody = strings.NewReader(`{ "order_status": "in_progress" }`)
	resp, err = http.NewRequest("PUT", tester.URL+"/api/order/"+orderIDStr+"/status", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	resp.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	r, err = client.Do(resp)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()

	if r.StatusCode != 201 {
		t.Fatalf("Expected status 201, got %d", r.StatusCode)
	}

	t.Logf("Order status: in_progress")

	// Водитель изменил статус заказа на completed
	requestBody = strings.NewReader(`{ "order_status": "completed" }`)
	resp, err = http.NewRequest("PUT", tester.URL+"/api/order/"+orderIDStr+"/status", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	resp.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	r, err = client.Do(resp)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()

	if r.StatusCode != 201 {
		t.Fatalf("Expected status 201, got %d", r.StatusCode)
	}

	t.Logf("Order status: completed")

	// Попытка отменить заказ
	res, err = http.Post(tester.URL+"/api/order/"+orderIDStr+"/cancel", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 422 {
		t.Fatalf("Expected status 422, got %d", res.StatusCode)
	}

	t.Logf("Orders with ID: %d can not be canceled", response.OrderId)
}

func TestPostDriverStatusSuccess(t *testing.T) {
	tester := httptest.NewServer(App())
	defer tester.Close()

	requestBody := strings.NewReader(`{"driver_id": 2, "available": true, "current_location": {"lat": 12.9343, "ing": 22.3061}}`)
	res, err := http.Post(tester.URL+"/api/driver/status", "application/json", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", res.StatusCode)
	}
}
