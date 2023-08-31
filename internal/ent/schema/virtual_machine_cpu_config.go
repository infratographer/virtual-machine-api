package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/x/gidx"
)

// VirtualMachineCPUConfig holds the schema definition for the VirtualMachineCPUConfig entity.
type VirtualMachineCPUConfig struct {
	ent.Schema
}

// Fields of the VirtualMachineCPUConfig.
func (VirtualMachineCPUConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Comment("The ID for the virtual machaine cpu config.").
			Annotations(
				entgql.OrderField("ID"),
			),
		field.Int("cores").
			Comment("The number of cores for this virtual machine.").
			Annotations(
				entgql.OrderField("cores"),
			),
		field.Int("sockets").
			Comment("The number of sockets for this virtual machine.").
			Annotations(
				entgql.OrderField("sockets"),
			),
	}
}

// Edges of the VirtualMachineCPUConfig.
func (VirtualMachineCPUConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("virtual_machine", VirtualMachine.Type).
			Unique().
			Immutable(),
	}
}
