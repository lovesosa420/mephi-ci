package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус код: ожидали %v, получили %v", http.StatusOK, status)
	}

	expected := "Hello"
	if rr.Body.String() != expected {
		t.Errorf("Неверный ответ: ожидали '%v', получили '%v'", expected, rr.Body.String())
	}
}
