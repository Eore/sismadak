package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/eore/sismadak/models"
)

//Pokja adalah objek balikan repositori untuk di gunakan pada handler
type Pokja struct {
	dbCon *sql.DB
}

//InisiasiRepoPokja adalah fungsi untuk menginisiasi repositori pokja
func InisiasiRepoPokja(koneksiDatabase *sql.DB) Pokja {
	return Pokja{
		dbCon: koneksiDatabase,
	}
}

//TambahPokja adalah fungsi repositori untuk menambah data pokja ke dalam database
func (r Pokja) TambahPokja(model models.Pokja) models.Pokja {
	var m models.Pokja
	row := r.dbCon.QueryRow(`
		INSERT INTO pokja
		VALUES ($1, $2)
		RETURNING *
	`,
		model.KodePokja,
		model.KeteranganPokja,
	)
	err := row.Scan(
		&m.KodePokja,
		&m.KeteranganPokja,
	)
	if err != nil {
		log.Printf("Gagal menambahkan data pokja : %+v\n", model)
	}
	return m
}

//ListPokja adalah fungsi repositori untuk menampilkan list data pokja dalam database
func (r Pokja) ListPokja() []models.Pokja {
	var m models.Pokja
	var list []models.Pokja
	rows, err := r.dbCon.Query("SELECT * FROM pokja")
	if err != nil {
		log.Println("Gagal menjalankan query")
		return list
	}
	for rows.Next() {
		err := rows.Scan(
			&m.KodePokja,
			&m.KeteranganPokja,
		)
		if err != nil {
			log.Println("Gagal membaca data")
			return list
		}
		list = append(list, m)
	}
	return list
}

//EditPokja adalah fungsi repositori untuk mengedit data pokja di dalam database
func (r Pokja) EditPokja(kodePokja string, data map[string]interface{}) models.Pokja {
	var m models.Pokja
	fields := []string{}
	values := []interface{}{kodePokja}
	i := 2
	for key := range data {
		fields = append(fields, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, data[key])
		i++
	}
	sqlStr := fmt.Sprintf(`
		UPDATE pokja 
		SET %s 
		WHERE kode_pokja = $1 
		RETURNING *
	`, strings.Join(fields, ", "))
	row := r.dbCon.QueryRow(sqlStr, values...)
	err := row.Scan(
		&m.KodePokja,
		&m.KeteranganPokja,
	)
	if err != nil {
		log.Printf("Gagal mengedit data pokja : %+v\n", data)
	}
	return m
}

//HapusPokja adalah fungsi repositori untuk menghapus data pokja dalam database sesuai kode pokja
func (r Pokja) HapusPokja(kodePokja string) models.Pokja {
	var m models.Pokja
	row := r.dbCon.QueryRow(`
		DELETE FROM pokja
		WHERE kode_pokja = $1
		RETURNING *
	`, kodePokja)
	err := row.Scan(
		&m.KodePokja,
		&m.KeteranganPokja,
	)
	if err != nil {
		log.Printf("Gagal menghapus data pokja : %+v\n", kodePokja)
	}
	return m
}
