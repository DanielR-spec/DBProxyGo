package main

import "fmt"
import "database/sql"
import _ "go-sql-driver/mysql"

db, err := sql.Open("mysql", "motivy_user:53745344-e5dc-4406-a163-d765c3bf5ecf@motivy-prod-postgress.cgyidm3l3vnq.us-east-1.rds.amazonaws.com/motivyprod?parseTime=true")


func main() {
    fmt.Println("Database conexion test")
    err := db.Ping()
}
