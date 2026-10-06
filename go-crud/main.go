package main

import (
	"go-crud/internal/database"
	"go-crud/internal/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errror Loading .env")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hi fiber",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid input",
			})
		}

		err := db.QueryRow(
			c.Context(),
			`INSERT INTO users(name, email) 
			 VALUES ($1, $2)
			 RETURNING id`,
			user.Name,
			user.Email,
		).Scan(&user.ID)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to create user",
			})
		}

		return c.Status(201).JSON(user)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := db.Query(
			c.Context(),
			`SELECT id,name,email FROM users;`,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to fetch users",
			})
		}

		defer rows.Close()

		var users []models.User

		for rows.Next() {
			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.Email,
				&user.Name,
			)

			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "failed to read user",
				})
			}

			users = append(users, user)

		}
		return c.JSON(users)
	})

	app.Get("users/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		var user models.User

		err := db.QueryRow(c.Context(),`SELECT id,name,email FROM users WHERE id = $1`,id).Scan(&user.ID,&user.Name,&user.Email)

		if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	return c.JSON(user)
	})

	log.Fatal(app.Listen(":3000"))
}
