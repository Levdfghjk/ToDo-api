package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"study/models"
	"study/storage"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var NewTask models.Task

	err := json.NewDecoder(r.Body).Decode(&NewTask)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	for _, task := range storage.Tasks {
		if task.Title == NewTask.Title {
			http.Error(w, "repeated title", http.StatusBadRequest)
			return
		}
	}

	if NewTask.Title == "" {
		http.Error(w, "empty title", http.StatusBadRequest)
		return
	}

	NewTask.ID = storage.NextID
	NewTask.Done = false

	storage.NextID++
	storage.Tasks = append(storage.Tasks, NewTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "problem with conversation", http.StatusBadRequest)
		return
	}

	for i, task := range storage.Tasks {
		if task.ID == id {
			storage.Tasks[i].Done = true

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(storage.Tasks[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		http.Error(w, "problems with conversation", http.StatusMethodNotAllowed)
		return
	}

	for i, task := range storage.Tasks {
		if task.ID == id {
			storage.Tasks = append(storage.Tasks[:i], storage.Tasks[i+1:]...)
			w.Write([]byte("successful delete"))
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func SearchByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "problems with conversation", http.StatusMethodNotAllowed)
		return
	}

	for i, task := range storage.Tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(storage.Tasks[i])

			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
