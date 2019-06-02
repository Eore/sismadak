package routes

import (
	"fmt"
	"net/http"

	"github.com/eore/sismadak/handlers"
)

func RouterPokja(router *http.ServeMux, handler handlers.Pokja) {
	router.HandleFunc("/pokja", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handler.ListPokja(w, r)
		case "POST":
			handler.TambahPokja(w, r)
		default:
			fmt.Fprintln(w, "method salah")
		}
	})
	router.HandleFunc("/pokja/", handler.ListPokja)
}
