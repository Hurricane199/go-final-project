package api

import (
	"encoding/json"
	"net/http"

	"github.com/Hurricane199/go-final-project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{
			"error": "не указан идентификатор",
		})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "не указан заголовок задачи",
		})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if !isRepeatValid(task.Repeat) {
		writeJSON(w, map[string]string{
			"error": "недопустимый формат повторения",
		})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, struct{}{})
}
