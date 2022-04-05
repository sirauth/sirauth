package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Realm holds the schema definition for the Realm entity.
type Realm struct {
	ent.Schema
}

// Fields of the Realm.
func (Realm) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Realm.
func (Realm) Edges() []ent.Edge {
	return nil
}
