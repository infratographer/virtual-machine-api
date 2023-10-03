package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

const (
	minVMMemory     = 1
	defaultVMMemory = 8
)

// VirtualMachineMemoryConfig holds the schema definition for the VirtualMachineMemoryConfig entity.
type VirtualMachineMemoryConfig struct {
	ent.Schema
}

// Fields of the VirtualMachineMemoryConfig.
func (VirtualMachineMemoryConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Comment("The ID for the virtual machine memory config.").
			Annotations(
				entgql.OrderField("ID"),
			),
		field.Int("size").
			Min(minVMMemory).
			Default(defaultVMMemory).
			Comment("The size of memory for this virtual machine.").
			Annotations(
				entgql.OrderField("SIZE"),
			),
	}
}

// Edges of the VirtualMachineMemoryConfig.
func (VirtualMachineMemoryConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("virtual_machine", VirtualMachine.Type).
			Unique().
			Immutable(),
	}
}

// Annotations of VirtualMachineMemoryConfig
func (VirtualMachineMemoryConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		entgql.Type("VirtualMachineMemoryConfig"),
		prefixIDDirective(VirtualMachinePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a virtual machine memory config."),
			entgql.MutationUpdate().Description("Input information to update a virtual machine memory config."),
		),
	}
}
