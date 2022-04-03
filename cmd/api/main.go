package main

import (
	"context"
	"fmt"
	"os"
	"sirauth/pkg/entities"
	"sirauth/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/elliotchance/pie/v2"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func main() {
	dsn := "host=localhost user=sirauth password=Password01 dbname=sirauth sslmode=disable"
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)

	engine := html.New("./templates", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	clientRoutes := routes.NewClientRoutes(dbpool)

	clientRoutes.AddRoutes(app)

	app.Get("/authorize", func(ctx *fiber.Ctx) error {
		clientId := ctx.Query("client_id")
		redirectUri := ctx.Query("redirect_uri")

		if clientId == "" {
			return ctx.Status(404).JSON(map[string]string{"error": "Invalid Client Id"})
		}

		if redirectUri == "" {
			return ctx.Status(404).JSON(map[string]string{"error": "Invalid Redirect URI"})
		}

		client, err := entities.GetByClientID(dbpool, clientId)
		if err != nil {
			return err
		}

		hasIt := pie.Contains(client.RedirectUrls, redirectUri)
		if !hasIt {
			return ctx.Status(401).JSON(map[string]string{"error": "Invalid Redirect URI"})
		}

		reqId := uuid.New().String()
		err = redisClient.Set(reqId, ctx.Request().URI().QueryString(), 0).Err()
		if err != nil {
			return err
		}

		return ctx.Render("approve", nil)
	})

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
