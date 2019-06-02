package drivers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //driver postgres
)

//KoneksiDatabase digunakan untuk membuat koneksi ke database
func KoneksiDatabase(driver, host, port, database, user, pass string) *sql.DB {
	var dbURL string
	dbURL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, user, pass, host, port, database)
	con, err := sql.Open(driver, dbURL)
	if err != nil {
		log.Fatalln("Koneksi ke database gagal")
	}
	if _, err := con.Exec("SELECT 1+1"); err != nil {
		con.Close()
		dbURL = fmt.Sprintf("%s://%s:%s@%s:%s?sslmode=disable", driver, user, pass, host, port)
		conTemp, err := sql.Open(driver, dbURL)
		log.Printf("Membuat database baru dengan nama %s\n", database)
		createDB := fmt.Sprintf("CREATE DATABASE %s", database)
		if _, err := conTemp.Query(createDB); err != nil {
			fmt.Println("Gagal membuat database")
		}
		conTemp.Close()
		dbURL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, user, pass, host, port, database)
		conNew, err := sql.Open(driver, dbURL)
		if err != nil {
			log.Fatalln("Koneksi ke database gagal")
		}
		return conNew
	}
	return con
}
