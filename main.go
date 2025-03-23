package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "log"
    "net/http"
    "os"
)

type Recommendation struct {
    MovieID int     `json:"movie_id"`
    Title   string  `json:"title"`
    Score   float64 `json:"score"`
}

func getRecommendations(c *gin.Context) {
//    userID := c.Param("user_id")
    recommendations := []Recommendation{
        {MovieID: 1, Title: "Inception", Score: 0.95},
        {MovieID: 2, Title: "Interstellar", Score: 0.92},
        }
        c.JSON(http.StatusOK, recommendations)
    // Логика получения рекомендаций (например, вызов Rust-сервиса)
}

func main() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Формируем строку подключения
    connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
        dbUser,
        dbPassword,
        dbHost,
        dbPort,
        dbName,
    )

    // Создайте пул соединений
    dbpool, err := pgxpool.New(context.Background(), connString)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v\n", err)
    }
    defer dbpool.Close()

    // Проверка соединения
    err = dbpool.Ping(context.Background())
    if err != nil {
        log.Fatalf("Unable to ping database: %v\n", err)
    }
    fmt.Println("Successfully connected to the database!")

    // Пример выполнения запроса
    var greeting string
    err = dbpool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
    if err != nil {
        log.Fatalf("QueryRow failed: %v\n", err)
    }
    fmt.Println(greeting)

    // Пример вставки данных
    _, err = dbpool.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
    if err != nil {
        log.Fatalf("Exec failed: %v\n", err)
    }
    fmt.Println("Data inserted successfully!")

    // Пример выборки данных
    var name, email string
    err = dbpool.QueryRow(context.Background(), "SELECT name, email FROM users WHERE id = $1", 1).Scan(&name, &email)
    if err != nil {
        log.Fatalf("QueryRow failed: %v\n", err)
    }
    fmt.Printf("User: %s, Email: %s\n", name, email)

    r := gin.Default()
    r.GET("/recommendations/:user_id", getRecommendations)
    r.Run(":8080")
}
