package main

import (
    "database/sql"
    "fmt"
    "net/url"

    _ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
  )


const (
    host     = "host"
    port     = 5432
    user     = "user"
    password = "password"
    dbname   = "dbname"
  )

func main() {

    // Encoding the password
    encodedPassword := url.QueryEscape(password)

    // Constructing the connection string
    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", user, encodedPassword, host, port, dbname)
    fmt.Println(connectionString)

    // Connect to the database
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        panic(err)
    }

    // Test the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    defer db.Close()
  
}
