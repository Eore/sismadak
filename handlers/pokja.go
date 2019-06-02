package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eore/sismadak/models"

	"github.com/eore/sismadak/repositories"
)

//Pokja adalah objek handler untuk digunakan di router
type Pokja struct {
	repo repositories.Pokja
}

//InisiasiPokjaHandler adalah fungsi untuk menginisiasi handler pokja
func InisiasiPokjaHandler(repo repositories.Pokja) Pokja {
	return Pokja{
		repo: repo,
	}
}

func (p Pokja) TambahPokja(w http.ResponseWriter, r *http.Request) {
	var m models.Pokja
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Fprintln(w, `"format data salah"`)
	}
	res := p.repo.TambahPokja(m)
	if (models.Pokja{}) == res {
		w.WriteHeader(400)
		fmt.Fprintln(w, `"gagal menambah data"`)
	}
	data, _ := json.Marshal(res)
	fmt.Fprintln(w, string(data))
}

func (p *Pokja) ListPokja(w http.ResponseWriter, r *http.Request) {
	list := p.repo.ListPokja()
	data, _ := json.Marshal(list)
	fmt.Fprintln(w, string(data))
}

func (p *Pokja) EditPokja(w http.ResponseWriter, r *http.Request) {
	kodePokja := r.FormValue("kode_pokja")
	if kodePokja == "" {
		fmt.Fprintln(w, `"kode pokja tidak boleh kosong"`)
		return
	}
	var dataBody map[string]interface{}
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &dataBody)
	if err != nil {
		fmt.Fprintln(w, `"format data salah"`)
	}
	res := p.repo.EditPokja(kodePokja, dataBody)
	if (models.Pokja{}) == res {
		w.WriteHeader(400)
		fmt.Fprintln(w, `"gagal mengedit data"`)
	}
	data, _ := json.Marshal(res)
	fmt.Fprintln(w, string(data))
}

func (p *Pokja) HapusPokja(w http.ResponseWriter, r *http.Request) {
	kodePokja := r.FormValue("kode_pokja")
	if kodePokja == "" {
		fmt.Fprintln(w, `"kode pokja tidak boleh kosong"`)
		return
	}
	res := p.repo.HapusPokja(kodePokja)
	if (models.Pokja{}) == res {
		w.WriteHeader(400)
		fmt.Fprintln(w, `"gagal menghapus data"`)
	}
	data, _ := json.Marshal(res)
	fmt.Fprintln(w, string(data))
}
