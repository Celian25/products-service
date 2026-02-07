package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FilterType holds the schema definition for the FilterType entity.
type FilterType struct {
	ent.Schema
}

// Fields of the FilterType.
func (FilterType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("tech_name").Optional().Unique(),
	}
}

// Edges of the FilterType.
func (FilterType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("filters", Filter.Type).Ref("filter_type"),
	}
}

func (FilterType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(),
	}
}
