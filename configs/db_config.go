package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadEnv(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	return nil
}

func SetupDB() (*sql.DB, error) {
	// Mendapatkan konfigurasi database dari environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Konfigurasi koneksi ke database MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Menguji koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}
