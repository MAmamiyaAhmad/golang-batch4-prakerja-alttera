package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Fungsi untuk mengambil semua buku dari database
func GetAllBooks(db *sql.DB) ([]Book, error) {
	books := make([]Book, 0)

	// Query ke database untuk mendapatkan semua buku
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Fungsi untuk mengambil buku berdasarkan ID dari database
func GetBookByID(db *sql.DB, id int) (*Book, error) {
	var book Book

	// Query ke database untuk mendapatkan buku berdasarkan ID
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Fungsi untuk menyimpan buku ke database
func CreateBook(db *sql.DB, book *Book) (int64, error) {
	// Query ke database untuk menyimpan buku
	result, err := db.Exec("INSERT INTO books (title, author, publisher) VALUES (?, ?, ?)", book.Title, book.Author, book.Publisher)
	if err != nil {
		return 0, err
	}

	// Mengambil ID buku yang baru saja disimpan
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Fungsi untuk memperbarui buku dengan ID yang diberikan di database
func UpdateBook(db *sql.DB, id int, book *Book) (int64, error) {
	// Query ke database untuk memperbarui buku
	result, err := db.Exec("UPDATE books SET title = ?, author = ?, publisher = ? WHERE id = ?", book.Title, book.Author, book.Publisher, id)
	if err != nil {
		return 0, err
	}

	// Mengambil jumlah baris yang terpengaruh oleh query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Fungsi untuk menghapus buku dengan ID yang diberikan dari database
func DeleteBook(db *sql.DB, id int) (int64, error) {
	// Query ke database untuk menghapus buku
	result, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	// Mengambil jumlah baris yang terpengaruh oleh query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
