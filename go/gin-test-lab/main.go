package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gin-test-lab/models"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const createBooks = `
CREATE TABLE books (
isbn char(14) NOT NULL,
title varchar(255) NOT NULL,
author varchar(255) NOT NULL,
price decimal(5,2) NOT NULL
	);`

const insertBooks = `
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);`

const alterBooks = `ALTER TABLE books ADD PRIMARY KEY (isbn);`

var dbPool *pgxpool.Pool

// func init() {
// 	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/bookstore"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	r, err := db.Query("SELECT * FROM books")
// 	if err != nil {
// 		fmt.Printf("INTI: error selecting books: %s", err)
// 	}
// 	defer r.Close()

// 	if len(r) == 0 {

// 		_, err = db.Query(createBooks)
// 		if err != nil {
// 			fmt.Printf("INIT: error creating data: %s", err)
// 		}

// 		_, err = db.Query(insertBooks)
// 		if err != nil {
// 			fmt.Printf("INIT: error inserting data: %s", err)
// 		}

// 		_, err = db.Query(alterBooks)
// 		if err != nil {
// 			fmt.Printf("INIT: error altering table: %s", err)
// 		}

// 	}

// }

type Env struct {
	// Replace the reference to models.BookModel with an interface
	// describing its methods instead. All the other code remains exactly
	// the same.
	books interface {
		All() ([]models.Book, error)
	}
	people interface {
		All() ([]any, error)
		Count() (any, error)
		Any() (any, error)
	}
}

func main() {
	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, "postgres://admin:root@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		books: models.BookModel{DB: dbPool},
		people: models.PeopleModel{DB: dbPool},
	}

	http.HandleFunc("/books", env.booksIndex)
	http.HandleFunc("/people", env.PeopleIndex)
	http.HandleFunc("/people/count", env.PeopleCount)
	http.ListenAndServe(":3000", nil)

	log.Print("listening on port 3000")
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := env.books.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	log.Print("handling books index")

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

func (env *Env) PeopleIndex(w http.ResponseWriter, r *http.Request) {
	ppl, err := env.people.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	log.Printf("%v\n", ppl)
	fmt.Fprintf(w, "%v\n", ppl)

	// for _, p := range ppl {
	// 	fmt.Fprintf(w, "%s, %s, %v\n", p.FirstName, p.LastName, p.Age)
	// }
}

func (env *Env) PeopleCount(w http.ResponseWriter, r *http.Request) {
	count, err := env.people.Count()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "There are %v people in the table", count)

}

func (env *Env) PeopleAny(w http.ResponseWriter, r *http.Request) {
	ppl, err := env.people.Any()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprint(w, ppl)
}
