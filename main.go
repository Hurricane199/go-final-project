package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "7540"
const webDir = "web"

func main() {
	fmt.Println("Сервер запускается ...")

	// Получаем порт из переменной окружения TODO_PORT
	// Если переменная не установлена, используем значение по умолчанию

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Println("Сервер запускается на порту:", port)

	// Инициализируем HTTP-роутер
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(webDir))
	mux.Handle("/", fileServer)

	// Запускаем HTTP-сервер
	fmt.Println("Сервер запущен на порту:", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
