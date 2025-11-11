package main

import (
    "log"
    "net/http"
    "os"

    "project1/internal/user"
    "project1/pkg/database"
)

func main() {
    log.Println("Starting Dating App...")

    db, err := database.NewSQLiteDB("dating_app.db")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            name VARCHAR(100) NOT NULL,
            age INTEGER NOT NULL CHECK (age >= 18 AND age <= 100),
            gender VARCHAR(10) NOT NULL CHECK (gender IN ('male', 'female')),
            description TEXT,
            looking_for VARCHAR(10) NOT NULL CHECK (looking_for IN ('male',
                                                        'female', 'both')),
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

    log.Println("Database initialized successfully")

    userRepo := user.NewRepository(db)
    userService := user.NewService(userRepo)
    userHandler := user.NewHandler(userService)

    http.Handle("/api/v1/users/", userHandler)

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed",
					   http.StatusMethodNotAllowed)
			return
		}
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "ok"}`))
    })

    port := getEnv("PORT", "8080")
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
