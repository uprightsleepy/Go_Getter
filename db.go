package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDB() (*sql.DB, error) {
	dbUser, err := GetSecret(DbUsername)
	if err != nil {
		log.Fatalf("Error fetching DB username: %v", err)
		return nil, err
	}

	dbPassword, err := GetSecret(DbPassword)
	if err != nil {
		log.Fatalf("Error fetching DB password: %v", err)
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, 5432, dbUser, dbPassword, DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}

func InsertBook(db *sql.DB, entry book) error {
	queryCheck := `SELECT COUNT(*) FROM books WHERE title = $1 AND img_url = $2`
	var count int
	err := db.QueryRow(queryCheck, entry.Title, entry.ImgUrl).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking if book exists: %w", err)
	}

	if count > 0 {
		log.Printf("Book already exists, skipping: %s", entry.Title)
		return nil
	}

	queryInsert := `INSERT INTO books (title, price, img_url, rating) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(queryInsert, entry.Title, entry.Price, entry.ImgUrl, entry.Rating)
	if err != nil {
		return fmt.Errorf("error inserting book: %w", err)
	}

	log.Printf("Inserted new book into database: %s", entry.Title)
	return nil
}
