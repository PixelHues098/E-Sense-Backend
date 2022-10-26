package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

// Env holds database connection to Postgres
type Env struct {
    DB *sql.DB
}

// database variables
// usually we should get them from env like os.Getenv("variableName")
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "root"
    dbname   = "esense"
)

// ConnectDB tries to connect DB and on succcesful it returns
// DB connection string and nil error, otherwise return empty DB and the corresponding error.
func ConnectDB() (*sql.DB, error) {
    connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", connString)
    if err != nil {
        log.Printf("failed to connect to database: %v", err)
        return &sql.DB{}, err
    } else {
        fmt.Println("Connection success")
    }
    return db, nil
}
