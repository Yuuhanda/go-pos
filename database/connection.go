package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    
    "your-app-name/config"
    
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Initialize sets up the database connection
func Initialize() {
    dbConfig := config.GetDBConfig()
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_general_ci",
        dbConfig.User,
        dbConfig.Password,
        dbConfig.Host,
        dbConfig.Port,
        dbConfig.DBName)
    
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }
    
    // Configure connection pool
    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(10)
    DB.SetConnMaxLifetime(5 * time.Minute)
    
    // Test connection
    if err := DB.Ping(); err != nil {
        log.Fatalf("Could not establish database connection: %v", err)
    }
    
    log.Println("Database connection established successfully!")
}

// Close closes the database connection
func Close() {
    if DB != nil {
        DB.Close()
        log.Println("Database connection closed")
    }
}
