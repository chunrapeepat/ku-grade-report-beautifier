package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/app"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("server is running on port", port)
	http.ListenAndServe(":"+port, mux)
}
