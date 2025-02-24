package main

import (
    "database/sql"
    "github.com/gofiber/fiber/v2"
    _ "github.com/lib/pq"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "password"
    DB_NAME     = "taskdb"
    DB_HOST     = "localhost"
    DB_PORT     = "5432"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME + "?sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, name TEXT NOT NULL)")
    if err != nil {
        panic(err)
    }
}

func getTasks(c *fiber.Ctx) error {
    rows, err := db.Query("SELECT id, name FROM tasks")
    if err != nil {
        return err
    }
    defer rows.Close()

    var tasks []map[string]interface{}
    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        tasks = append(tasks, map[string]interface{}{"id": id, "name": name})
    }

    return c.JSON(tasks)
}

func addTask(c *fiber.Ctx) error {
    var task map[string]string
    if err := c.BodyParser(&task); err != nil {
        return err
    }

    _, err := db.Exec("INSERT INTO tasks (name) VALUES ($1)", task["name"])
    if err != nil {
        return err
    }

    return c.JSON(task)
}

func main() {
    initDB()

    app := fiber.New()

    app.Get("/tasks", getTasks)
    app.Post("/tasks", addTask)

    app.Listen(":3000")
}
