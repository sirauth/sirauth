package realm

import (
	"context"
	"log"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"
)

const TableName = "realms"

type Realm struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var orm = sqlbuilder.NewStruct(new(Realm))

func New() Realm {
	return Realm{}
}

func GetById(pool *pgxpool.Pool, id int) (*Realm, error) {
	entity := Realm{}
	ub := orm.SelectFrom(TableName)
	ub.Where(ub.E("id", id))
	sql, args := ub.Build()

	err := pool.QueryRow(context.Background(),
		sql,
		args...).Scan(&entity.Id, &entity.Name)
	if err != nil {
		log.Printf("Select on %s for id %d failed: %v", TableName, id, err)
		return nil, err
	}

	return &entity, nil
}

func Update(pool *pgxpool.Pool, entity *Realm) error {
	ub := orm.Update(TableName, entity)
	ub.Where(ub.E("id", entity.Id))
	sql, args := ub.Build()

	_, err := pool.Exec(context.Background(),
		sql,
		args...)
	if err != nil {
		log.Printf("Update on %s failed: %v", TableName, err)
		return err
	}

	return nil
}

func Create(pool *pgxpool.Pool, entity *Realm) error {
	ub := orm.InsertInto(TableName, entity)
	sql, args := ub.Build()

	sql = sql + "; RETURNING id, name"

	err := pool.QueryRow(context.Background(),
		sql,
		args...).Scan(&entity.Id, &entity.Name)
	if err != nil {
		log.Printf("Insert on %s failed: %v", TableName, err)
		return err
	}

	return nil
}
