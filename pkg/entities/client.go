package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	Id           int      `json:"id"`
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	RedirectUrls []string `json:"redirectUrls"`
}

var clientStruct = sqlbuilder.NewStruct(new(Client))

func GetClientByClientId(pool *pgxpool.Pool, id string) (Client, error) {
	var client Client
	err := pool.QueryRow(context.Background(),
		"SELECT id, client_id, client_secret, redirect_urls FROM clients WHERE client_id = $1",
		id).
		Scan(&client.Id, &client.ClientID, &client.ClientSecret, &client.RedirectUrls)
	if err != nil {
		fmt.Errorf("QueryRow failed: %v\n", err)
		return client, err
	}

	return client, nil
}

func UpdateClient(pool *pgxpool.Pool, client *Client) error {
	ub := clientStruct.Update("clients", client)
	ub.Where(ub.E("id", client.Id))
	sql, args := ub.Build()

	_, err := pool.Exec(context.Background(),
		sql,
		args...)
	if err != nil {
		log.Printf("QueryRow failed: %v", err)
		return err
	}

	return nil
}
