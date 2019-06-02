package handlers

import (
	"fmt"
	"net/http"

	"github.com/eore/sismadak/repositories"
)

type Pokja struct {
	r repositories.Pokja
}

func InisiasiPokjaHandler(repo repositories.Pokja) Pokja {
	return Pokja{
		r: repo,
	}
}

func (p Pokja) TambahPokja(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "tambah")
}

func (p *Pokja) ListPokja(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list")
}
