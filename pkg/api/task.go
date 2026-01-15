package api

import (
	"net/http"

	"github.com/Hurricane199/go-final-project/pkg/db"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)

	case http.MethodGet:
		getTaskHandler(w, r)

	case http.MethodPut:
		updateTaskHandler(w, r)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			writeJSON(w, map[string]string{
				"error": "не указан идентификатор",
			})
			return
		}

		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, struct{}{})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, map[string]string{
			"error": "метод недоступен",
		})
	}
}
