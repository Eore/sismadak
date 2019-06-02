package repositories

import (
	"database/sql"
	"log"

	"github.com/eore/sismadak/models"
)

type Pokja struct {
	dbCon *sql.DB
}

func InisiasiRepoPokja(koneksiDatabase *sql.DB) Pokja {
	return Pokja{
		dbCon: koneksiDatabase,
	}
}

func (r Pokja) TambahPokja(model models.Pokja) models.Pokja {
	rows := r.dbCon.QueryRow(`
		INSERT INTO pokja
		VALUES ($1, $2)
	`,
		model.KodePokja,
		model.KeteranganPokja,
	)
	var ret models.Pokja
	err := rows.Scan(
		&ret.KodePokja,
		&ret.KeteranganPokja,
	)
	if err != nil {
		log.Printf("Gagal menambahkan data pokja : %+v\n", model)
	}
	return ret
}
