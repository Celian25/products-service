package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Text("description").Optional(),
		field.String("slug").NotEmpty(),
		field.String("photo_url").Optional(),
		field.Time("created_at").Default(time.Now).Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput,
				entgql.SkipMutationUpdateInput,
			)),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type).Unique().Required(),
		edge.To("filters", Filter.Type),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
	}
}
