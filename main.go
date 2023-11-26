package main

import "fmt"
import "database/sql"
import _ "go-sql-driver/mysql"

db, err := sql.Open("mysql", "")


func main() {
    fmt.Println("Database conexion test")
    err := db.Ping()
}
