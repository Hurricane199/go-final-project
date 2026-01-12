package api

import (
	"net/http"

	"github.com/Hurricane199/go-final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
