package handlers

import (
	"encoding/json"
	"net/http"
	"study/storage"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Tasks)
}
