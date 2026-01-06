package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Hurricane199/go-final-project/pkg/db"
	"github.com/Hurricane199/go-final-project/pkg/server"
)

const defaultPort = "7540"

const webDir = "web"

func main() {
	dbFile := os.Getenv("TODO_DBFILE")

	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	fmt.Println("Сервер запускается ...")

	// Получаем порт из переменной окружения TODO_PORT
	// Если переменная не установлена, используем значение по умолчанию

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Сервер запустился на порту: %s\n", port)
	log.Fatal(server.Run(port, webDir))

}
