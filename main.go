package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	_ "github.com/lib/pq"
)

func getDB() *sql.DB {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		const (
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "pass"
			dbname   = "postgres"
		)
		databaseURL = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func main() {
	db := getDB()
	fmt.Println(db.Ping())
	engine := html.New("./templates", ".html")
	engine.Reload(true)
	engine.Debug(true)
	engine.Layout("embed")
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", mainPage)

	log.Fatal(app.Listen(":8000"))
}

func mainPage(c *fiber.Ctx) error {
	data := fiber.Map{
		"Title": "CrossGen",
	}
	return c.Render("index", data, "layouts/main")
}
