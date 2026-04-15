package main

import (
	"fmt"
	"net/http"
	"study/handlers"

	"github.com/gorilla/mux"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	rounter := mux.NewRouter()

	rounter.HandleFunc("/ping", pingHandler).Methods("GET")
	rounter.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	rounter.HandleFunc("/tasks/{id}", handlers.SearchByID).Methods("GET")
	rounter.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	rounter.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	rounter.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	err := http.ListenAndServe(":9091", rounter)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
	}

	fmt.Println("Сервер запущен!")
}
