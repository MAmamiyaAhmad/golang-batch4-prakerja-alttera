package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"yourproject/configs"
	"yourproject/routes"
)

func main() {
	// Memuat konfigurasi dari file .env
	err := configs.LoadEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Mengatur koneksi ke database
	db, err := configs.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Menambahkan rute-rute aplikasi dan menyediakan koneksi ke database
	routes.RegisterRoutes(r, db)

	// Menentukan alamat dan port yang akan digunakan
	addr := "localhost:8080"

	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
