package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // Import the PostgreSQL driver
)

func Connect() *sql.DB {
    connStr := "user=username dbname=yourdbname sslmode=disable" // Update with your DB credentials
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    return db
}