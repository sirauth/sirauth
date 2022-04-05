package main

import (
	"fmt"
	"log"
	"sirauth/ent"
	"sirauth/pkg/entities/realm"
	"sirauth/pkg/routes"

	"database/sql"

	"entgo.io/ent/dialect"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open(dsn string) *ent.Client {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	dsn := "host=localhost user=sirauth password=Password01 dbname=sirauth sslmode=disable"
	dbClient := Open(dsn)

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

	clientRoutes := routes.NewClientRoutes(dbClient)
	realmRoutes := realm.NewRoutes(dbClient)

	clientRoutes.AddRoutes(app)
	realmRoutes.AddRoutes(app)
	/*
		app.Get("/authorize", func(ctx *fiber.Ctx) error {
			clientId := ctx.Query("client_id")
			redirectUri := ctx.Query("redirect_uri")

			if clientId == "" {
				return ctx.Status(404).JSON(map[string]string{"error": "Invalid Client Id"})
			}

			if redirectUri == "" {
				return ctx.Status(404).JSON(map[string]string{"error": "Invalid Redirect URI"})
			}

			client, err := entities.GetClientByClientId(dbpool, clientId)
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
	*/
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
