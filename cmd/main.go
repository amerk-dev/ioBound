package main

import (
	"fmt"
	"ioBound/internal/handler"
	"net/http"
)

func main() {
	server := handler.NewServer()

	http.HandleFunc("/task", server.CreateTask)
	http.HandleFunc("/task/", server.CheckStatus)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
