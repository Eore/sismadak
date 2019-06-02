package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eore/sismadak/drivers"
	"github.com/eore/sismadak/handlers"
	"github.com/eore/sismadak/helpers"
	"github.com/eore/sismadak/models"
	"github.com/eore/sismadak/repositories"
	"github.com/eore/sismadak/routes"
)

const (
	portServer = ":9000"
)

func inisiasiRouter(con *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	repoPokja := repositories.InisiasiRepoPokja(con)
	handlerPokja := handlers.InisiasiPokjaHandler(repoPokja)
	routes.RouterPokja(router, handlerPokja)

	return router
}

func inisiasiModel(con *sql.DB) {
	log.Println("Inisiasi model...")
	helpers.InisiasiModel(con, "pokja", models.Pokja{})
	log.Println("Inisiasi model selesai")
}

func main() {
	var host, port, database, user, pass string
	flag.StringVar(&host, "dbhost", "", "host database")
	flag.StringVar(&port, "dbport", "", "port database")
	flag.StringVar(&database, "database", "", "nama database")
	flag.StringVar(&user, "dbusername", "", "username database")
	flag.StringVar(&pass, "dbpassword", "", "password database")
	flag.Parse()

	if host == "" || port == "" || database == "" || user == "" || pass == "" {
		fmt.Fprintln(os.Stdout, "SISMADAK v1.0.0")
		fmt.Fprint(os.Stdout, "Input host database: ")
		fmt.Scan(&host)
		fmt.Fprint(os.Stdout, "Input port database: ")
		fmt.Scan(&port)
		fmt.Fprint(os.Stdout, "Input nama database: ")
		fmt.Scan(&database)
		fmt.Fprint(os.Stdout, "Input username database: ")
		fmt.Scan(&user)
		fmt.Fprint(os.Stdout, "Input password database: ")
		fmt.Scan(&pass)
	}

	log.Println("Koneksi ke database...")
	con := drivers.KoneksiDatabase("postgres", host, port, database, user, pass)
	log.Println("Koneksi ke database berhasil")

	inisiasiModel(con)
	router := inisiasiRouter(con)

	log.Printf("Server berjalan di port%s", portServer)
	err := http.ListenAndServe(portServer, router)
	if err != nil {
		log.Fatalln("Gagal memulai server")
	}
}
