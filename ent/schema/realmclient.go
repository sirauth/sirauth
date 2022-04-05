package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RealmOAuthClient holds the schema definition for the RealmOAuthClient entity.
type RealmOAuthClient struct {
	ent.Schema
}

func (RealmOAuthClient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "clients"},
	}
}

// Fields of the RealmOAuthClient.
func (RealmOAuthClient) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("realm_id"),
		field.String("client_secret"),
		field.String("client_id"),
		field.JSON("redirect_urls", []string{}),
	}
}

// Edges of the RealmOAuthClient.
func (RealmOAuthClient) Edges() []ent.Edge {
	return nil
}
