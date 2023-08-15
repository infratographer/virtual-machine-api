package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/vektah/gqlparser/v2/ast"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/virtual-machine-api/x/pubsubinfo"
)

// VirtualMachine holds the schema definition for the VM entity
type VirtualMachine struct {
	ent.Schema
}

// Mixin of VirtualMachine
func (VirtualMachine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of VirtualMachine
func (VirtualMachine) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the VirtualMachine.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(VirtualMachinePrefix) }),
		field.String("name").
			NotEmpty().
			Comment("The name of the Virtual Machine.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Comment("The ID for the owner of this Virtual Machine.").
			Immutable().
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("OWNER"),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("location_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the location of this virtual machine.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("userdata").
			Comment("The userdata for this virtual machine.").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
	}
}

// Edges of VirtualMachine
func (VirtualMachine) Edges() []ent.Edge {
	return nil
}

// Indexes of VirtualMachine
func (VirtualMachine) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations of VirtualMachine
func (VirtualMachine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a virtual machine on the graph."),
		prefixIDDirective(VirtualMachinePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new virtual machine."),
			entgql.MutationUpdate().Description("Update an existing virtual machine."),
		),
		pubsubinfo.Annotation{},
		// entx.EventsHookSubjectName("virtual-machine"),
	}
}

func prefixIDDirective(prefix string) entgql.Annotation {
	var args []*ast.Argument
	if prefix != "" {
		args = append(args, &ast.Argument{
			Name: "prefix",
			Value: &ast.Value{
				Raw:  prefix,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.Directives(entgql.NewDirective("prefixedID", args...))
}
