package feature_sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Book struct {
	Id            int        `db:"id"`
	Name          string     `db:"name"`
	Creator       string     `db:"creator"`
	Review        *string    `db:"review"`
	CreatedYear   int        `db:"created_year"`
	Compleated    *bool      `db:"compleated"`
	InLibraryTime time.Time  `db:"in_library_time"`
	CompleatedAt  *time.Time `db:"compleated_at"`
}

func CreateBooks(ctx context.Context, conn *pgx.Conn) error {
	sqlQuerry := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	name VARCHAR(200) NOT NULL,
	creator VARCHAR(200) NOT NULL,
	review VARCHAR(1000),
	created_year INT NOT NULL,
	compleated BOOLEAN,
	in_library_time TIMESTAMP NOT NULL,
	compleated_at TIMESTAMP
	);
	`
	_, err := conn.Exec(ctx, sqlQuerry)
	return err

}

func InsertBook(book *Book, ctx context.Context, conn *pgx.Conn) error {
	err := conn.QueryRow(ctx,
		`INSERT INTO books (name, creator, review, created_year, compleated, in_library_time, compleated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;`,
		book.Name, book.Creator, book.Review, book.CreatedYear, book.Compleated, book.InLibraryTime, book.CompleatedAt).Scan(&book.Id)
	return err
}

func GetBooks(ctx context.Context, conn *pgx.Conn) ([]Book, error) {
	rows, err := conn.Query(ctx, "SELECT * FROM books")
	if err != nil {
		return []Book{}, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		rows.Scan(&book.Id, &book.Name, &book.Creator, &book.Review, &book.CreatedYear, &book.Compleated, &book.InLibraryTime, &book.CompleatedAt)
		books = append(books, book)
	}
	return books, nil
}

func UpdateBook(book *Book, ctx context.Context, conn *pgx.Conn) error {
	sqlQuerry := `
	UPDATE books
	SET name=$2, creator=$3, review=$4, created_year=$5, compleated=$6, in_library_time=$7, compleated_at=$8
	WHERE id=$1
	`
	_, err := conn.Exec(ctx, sqlQuerry, book.Id, book.Name, book.Creator, book.Review, book.CreatedYear, book.Compleated, book.InLibraryTime, book.CompleatedAt)
	return err
}

func DeleteBooks(ids []int, ctx context.Context, conn *pgx.Conn) error {
	sqlQuerry := `
	DELETE FROM books WHERE id = ANY($1)
	`
	_, err := conn.Exec(ctx, sqlQuerry, ids)
	return err
}

func GetPages(ctx context.Context, conn *pgx.Conn, n int) error {
	offset := 0
	num := 0
	sqlQuerry := `
	SELECT * FROM books ORDER BY id ASC LIMIT $1 OFFSET $2
	`
	for {
		rows, err := conn.Query(ctx, sqlQuerry, n, offset)
		if err != nil {
			return err
		}
		var books []Book
		for rows.Next() {
			var book Book
			rows.Scan(&book.Id, &book.Name, &book.Creator, &book.Review, &book.CreatedYear, &book.Compleated, &book.InLibraryTime, &book.CompleatedAt)
			books = append(books, book)
		}
		num += 1
		fmt.Printf("Страница %d: %+v\n", num, books)
		if len(books) < n {
			break
		}
		offset += n
	}
	return nil

}
