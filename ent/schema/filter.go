package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Filter holds the schema definition for the Filter entity.
type Filter struct {
	ent.Schema
}

// Fields of the Filter.
func (Filter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Filter.
func (Filter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("filter_type", FilterType.Type).Unique().Required(),
		edge.From("products", Product.Type).Ref("filters"),
	}
}

func (Filter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(),
	}
}
