package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func main() {
	godotenv.Load()
	DATABASE_URL := os.Getenv("DATABASE_URL")

	conn, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fetchAllAlbums(conn)
	fetchAlbum(9, conn)

	albums := []Album{
		{Title: "Still Bill", Artist: "Bill Withers", Price: 56.99},
		{Title: "Black on Both Sides", Artist: "Mos Def", Price: 32.99},
		{Title: "Capital Punishment", Artist: "Big Pun", Price: 19.99},
		{Title: "The Low End Theory", Artist: "A Tribe Called Quest", Price: 29.99},
		{Title: "The Shape of Jazz to Come", Artist: "Ornette Coleman", Price: 39.99},
		{Title: "The Blueprint", Artist: "Jay-Z", Price: 24.99},
	}

	for _, album := range albums {
		fmt.Println(addAlbum(album, conn))
	}
}

func fetchAllAlbums(conn *pgxpool.Pool) {
	var albums []Album
	rows, err := conn.Query(context.Background(), "SELECT * FROM album") //.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			fmt.Println(err)
		}
		albums = append(albums, alb)
	}

	fmt.Println(albums)
}

func fetchAlbum(id int, conn *pgxpool.Pool) {
	var album = Album{}
	err := conn.QueryRow(context.Background(), "SELECT * FROM album WHERE id = $1", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(album)
}

func addAlbum(alb Album, conn *pgxpool.Pool) (int64, error) {
	result, err := conn.Exec(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("unable to execute the query. %v", err)
	}

	return result.RowsAffected(), nil
}
