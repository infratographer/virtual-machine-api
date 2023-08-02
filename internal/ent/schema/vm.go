package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

type VirtM struct {
	ent.Schema
}

func (VirtM) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

func (VirtM) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the VirtM.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(VirtMPrefix) }),
		field.Text("hostname").
			NotEmpty().
			Comment("The name of the VirtM.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("String"),
				entgql.OrderField("HOSTNAME"),
			),
	}
}

func (VirtM) Edges() []ent.Edge {
	return nil
}

func (VirtM) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hostname"), /* XXX think this is wrong -ians */
	}
}

func (VirtM) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a virtual machine on the graph."),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new virtual machine."),
			entgql.MutationUpdate().Description("Update an existing virtual machine."), /* XXX */
		),
	}
}
