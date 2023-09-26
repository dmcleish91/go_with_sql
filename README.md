# SQL Operations in main.go

This file contains an example of how to perform basic SQL operations in Go using the `pgx` library. The example code connects to a PostgreSQL database and performs the following operations:

- Fetch all albums from the `album` table
- Fetch a single album from the `album` table
- Insert new albums into the `album` table

## Connecting to the Database

The first step in performing SQL operations is to connect to the database. This is done using the `pgxpool.New` function, which takes a `context.Context` and a connection string as arguments. The connection string should be in the format `postgresql://user:password@host:port/database`.

```go
DATABASE_URL := "postgresql://user:password@host:port/database"
conn, err := pgxpool.New(context.Background(), DATABASE_URL)
if err != nil {
    // Handle error
}
defer conn.Close()
```

## Fetching Data

To fetch data from the database, you can use the `conn.Query` function. This function takes a `context.Context` and a SQL query string as arguments, and returns a `pgx.Rows` object that can be used to iterate over the results.

```go
rows, err := conn.Query(context.Background(), "SELECT * FROM album")
if err != nil {
    // Handle error
}
defer rows.Close()

for rows.Next() {
    var album Album
    err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
    if err != nil {
        // Handle error
    }
    // Do something with album
}
```

To fetch a single row from the database, you can use the `conn.QueryRow` function. This function takes a `context.Context` and a SQL query string as arguments, and returns a `pgx.Row` object that can be used to scan the result.

```go
var album Album
err := conn.QueryRow(context.Background(), "SELECT * FROM album WHERE id = $1", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
if err != nil {
    // Handle error
}
// Do something with album
```

## Inserting Data

To insert data into the database, you can use the `conn.Exec` function. This function takes a `context.Context` and a SQL query string as arguments, and returns a `pgx.CommandTag` object that can be used to check the number of rows affected.

```go
result, err := conn.Exec(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", album.Title, album.Artist, album.Price)
if err != nil {
    // Handle error
}
if result.RowsAffected() != 1 {
    // Handle error
}
```

That's it! This should give you a basic understanding of how to perform SQL operations in Go using the `pgx` library.