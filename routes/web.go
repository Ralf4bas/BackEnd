package routes

import (
	"fmt"
	"loginsystem/handlers"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/upload", handlers.UploadHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error menjalankan server:", err)
	}
}
