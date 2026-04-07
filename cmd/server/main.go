package main

import (
	"green-api-task/internal/client"
	"green-api-task/internal/handler"
	"log"
	"net/http"
)

func main() {
	greenClient := client.NewGreenClient()
	h := &handler.Handler{Client: greenClient}

	// Создаем новый роутер (Mux)
	mux := http.NewServeMux()

	// 1. Сначала регистрируем API.
	// Мы явно указываем обработчик для прокси
	mux.HandleFunc("/api/method", h.ProxyHandler)

	// 2. Регистрируем статику.
	// Важно: файловый сервер должен идти после специфичных API роутов
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	log.Println("🚀 Server running on http://localhost:8080")

	// Передаем mux в ListenAndServe вместо nil
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}