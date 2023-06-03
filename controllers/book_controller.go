package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Publisher string `json:"publisher"`
}

var db *sql.DB

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Mengambil semua buku dari database
	books := make([]Book, 0)

	// Query ke database untuk mendapatkan semua buku
	// db.Query("SELECT * FROM books")

	// Contoh sederhana untuk menambahkan 3 buku
	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1", Publisher: "Publisher 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2", Publisher: "Publisher 2"})
	books = append(books, Book{ID: 3, Title: "Book 3", Author: "Author 3", Publisher: "Publisher 3"})

	// Mengirimkan respon JSON dengan buku-buku
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	// Mengambil buku berdasarkan ID dari database
	// db.Query("SELECT * FROM books WHERE id = ?", bookID)

	// Contoh sederhana untuk mengambil buku dengan ID yang diberikan
	book := Book{ID: 1, Title: "Book 1", Author: "Author 1", Publisher: "Publisher 1"}

	// Mengirimkan respon JSON dengan buku yang ditemukan
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	// Membaca data buku yang dikirimkan dari body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Menyimpan buku ke database
	// db.Exec("INSERT INTO books (title, author, publisher) VALUES (?, ?, ?)", book.Title, book.Author, book.Publisher)

	log.Println("Book created:", book)

	// Mengirimkan respon JSON dengan buku yang berhasil dibuat
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	var book Book

	// Membaca data buku yang dikirimkan dari body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Memperbarui buku dengan ID yang diberikan di database
	// db.Exec("UPDATE books SET title = ?, author = ?, publisher = ? WHERE id = ?", book.Title, book.Author, book.Publisher, bookID)

	log.Println("Book updated:", book)

	// Mengirimkan respon JSON dengan buku yang berhasil diperbarui
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	// Menghapus buku dengan ID yang diberikan dari database
	// db.Exec("DELETE FROM books WHERE id = ?", bookID)

	log.Println("Book deleted:", bookID)

	// Mengirimkan respon kosong dengan status sukses
	w.WriteHeader(http.StatusOK)
}
