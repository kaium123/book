package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{

		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("title").
			Default("").Optional(),
		field.String("author").
			Default("").Optional(),
		field.Time("publication_year").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("is_deleted").
			Optional().
			Default(false),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
