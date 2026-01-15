package server

import (
	"net/http"

	"github.com/Hurricane199/go-final-project/pkg/api"
)

func Run(port, webDir string) error {
	// Инициализируем HTTP-роутер
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(webDir))
	mux.Handle("/", fileServer)

	// Инициализируем API-обработчики
	api.Init(mux)

	// Запускаем HTTP-сервер
	return http.ListenAndServe(":"+port, mux)
}
