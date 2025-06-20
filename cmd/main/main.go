package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoaoRafa19/portifolio/view"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := view.Home()

		component.Render(r.Context(), w)
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
