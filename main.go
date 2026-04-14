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

	err := http.ListenAndServe(":9091", rounter)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
