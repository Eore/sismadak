package models

/*
Pokja adalah model struktur untuk data tentang pokja.
KodePokja	berisi kode pokja (contoh: SKP, PMKP, dst).
KeteranganPokja berisi keterangan tentang pokja (contoh: Sasaran Keselamatan Pasien).
*/
type Pokja struct {
	KodePokja       string `json:"kode_pokja" sql:"VARCHAR(10) NOT NULL PRIMARY KEY"`
	KeteranganPokja string `json:"keterangan_pokja" sql:"TEXT"`
}
