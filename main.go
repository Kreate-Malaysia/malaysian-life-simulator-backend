package main
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
func main() {
    connStr := "postgresql://[user]:[password]@[neon_hostname]/[dbname]?sslmode=require"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    var version string
    if err := db.QueryRow("select version()").Scan(&version); err != nil {
        panic(err)
    }
    fmt.Printf("version=%s\n", version)
}