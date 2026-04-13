package main

import (
	feature_sql "Bases/feature_postgres"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {
	link := "postgres://postgres:Troitskiy2007@localhost:5432/postgres"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, link)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	err = feature_sql.CreateBooks(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("База создана")

	// for i := 1000; i < 1060; i++ {
	// 	rew := "Это было превосходно"
	// 	book := feature_sql.Book{Name: "Мстители", Creator: "Пушкин", Review: &rew, CreatedYear: i, InLibraryTime: time.Now()}
	// 	err = feature_sql.InsertBook(&book, ctx, conn)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Книга создана")
	// }

	// rew := "Это было УЖАСНО"
	// comp := true
	// book := feature_sql.Book{Id: 6, Name: "Мстители 5", Creator: "Пушкин", Review: &rew, CreatedYear: 2010, InLibraryTime: time.Now(), Compleated: &comp}
	// err = feature_sql.UpdateBook(&book, ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// ids := []int{4, 6}
	// err = feature_sql.DeleteBooks(ids, ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// books, err := feature_sql.GetBooks(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }
	// pp.Println(books)

	err = feature_sql.GetPages(ctx, conn, 11)
	if err != nil {
		panic(err)
	}

}
