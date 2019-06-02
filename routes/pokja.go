package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eore/sismadak/handlers"
)

func RouterPokja(router *http.ServeMux, handler handlers.Pokja) {
	router.HandleFunc("/pokja", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.Method, r.RemoteAddr)
		switch r.Method {
		case "GET":
			handler.ListPokja(w, r)
		case "POST":
			handler.TambahPokja(w, r)
		case "PUT":
			handler.EditPokja(w, r)
		default:
			w.WriteHeader(404)
			fmt.Fprintln(w, `"method salah"`)
		}
	})
}
