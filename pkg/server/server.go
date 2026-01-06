package server

import (
	"net/http"
)

func Run(port, webDir string) error {
	// Инициализируем HTTP-роутер
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(webDir))
	mux.Handle("/", fileServer)

	// Запускаем HTTP-сервер
	return http.ListenAndServe(":"+port, mux)
}
