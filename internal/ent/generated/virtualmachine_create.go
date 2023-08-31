// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinecpuconfig"
	"go.infratographer.com/x/gidx"
)

// VirtualMachineCreate is the builder for creating a VirtualMachine entity.
type VirtualMachineCreate struct {
	config
	mutation *VirtualMachineMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vmc *VirtualMachineCreate) SetCreatedAt(t time.Time) *VirtualMachineCreate {
	vmc.mutation.SetCreatedAt(t)
	return vmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vmc *VirtualMachineCreate) SetNillableCreatedAt(t *time.Time) *VirtualMachineCreate {
	if t != nil {
		vmc.SetCreatedAt(*t)
	}
	return vmc
}

// SetUpdatedAt sets the "updated_at" field.
func (vmc *VirtualMachineCreate) SetUpdatedAt(t time.Time) *VirtualMachineCreate {
	vmc.mutation.SetUpdatedAt(t)
	return vmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vmc *VirtualMachineCreate) SetNillableUpdatedAt(t *time.Time) *VirtualMachineCreate {
	if t != nil {
		vmc.SetUpdatedAt(*t)
	}
	return vmc
}

// SetName sets the "name" field.
func (vmc *VirtualMachineCreate) SetName(s string) *VirtualMachineCreate {
	vmc.mutation.SetName(s)
	return vmc
}

// SetOwnerID sets the "owner_id" field.
func (vmc *VirtualMachineCreate) SetOwnerID(gi gidx.PrefixedID) *VirtualMachineCreate {
	vmc.mutation.SetOwnerID(gi)
	return vmc
}

// SetLocationID sets the "location_id" field.
func (vmc *VirtualMachineCreate) SetLocationID(gi gidx.PrefixedID) *VirtualMachineCreate {
	vmc.mutation.SetLocationID(gi)
	return vmc
}

// SetUserdata sets the "userdata" field.
func (vmc *VirtualMachineCreate) SetUserdata(s string) *VirtualMachineCreate {
	vmc.mutation.SetUserdata(s)
	return vmc
}

// SetNillableUserdata sets the "userdata" field if the given value is not nil.
func (vmc *VirtualMachineCreate) SetNillableUserdata(s *string) *VirtualMachineCreate {
	if s != nil {
		vmc.SetUserdata(*s)
	}
	return vmc
}

// SetVMCPUConfigID sets the "vm_cpu_config_id" field.
func (vmc *VirtualMachineCreate) SetVMCPUConfigID(gi gidx.PrefixedID) *VirtualMachineCreate {
	vmc.mutation.SetVMCPUConfigID(gi)
	return vmc
}

// SetID sets the "id" field.
func (vmc *VirtualMachineCreate) SetID(gi gidx.PrefixedID) *VirtualMachineCreate {
	vmc.mutation.SetID(gi)
	return vmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vmc *VirtualMachineCreate) SetNillableID(gi *gidx.PrefixedID) *VirtualMachineCreate {
	if gi != nil {
		vmc.SetID(*gi)
	}
	return vmc
}

// SetVirtualMachineCPUConfigID sets the "virtual_machine_cpu_config" edge to the VirtualMachineCPUConfig entity by ID.
func (vmc *VirtualMachineCreate) SetVirtualMachineCPUConfigID(id gidx.PrefixedID) *VirtualMachineCreate {
	vmc.mutation.SetVirtualMachineCPUConfigID(id)
	return vmc
}

// SetVirtualMachineCPUConfig sets the "virtual_machine_cpu_config" edge to the VirtualMachineCPUConfig entity.
func (vmc *VirtualMachineCreate) SetVirtualMachineCPUConfig(v *VirtualMachineCPUConfig) *VirtualMachineCreate {
	return vmc.SetVirtualMachineCPUConfigID(v.ID)
}

// Mutation returns the VirtualMachineMutation object of the builder.
func (vmc *VirtualMachineCreate) Mutation() *VirtualMachineMutation {
	return vmc.mutation
}

// Save creates the VirtualMachine in the database.
func (vmc *VirtualMachineCreate) Save(ctx context.Context) (*VirtualMachine, error) {
	vmc.defaults()
	return withHooks(ctx, vmc.sqlSave, vmc.mutation, vmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vmc *VirtualMachineCreate) SaveX(ctx context.Context) *VirtualMachine {
	v, err := vmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmc *VirtualMachineCreate) Exec(ctx context.Context) error {
	_, err := vmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmc *VirtualMachineCreate) ExecX(ctx context.Context) {
	if err := vmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vmc *VirtualMachineCreate) defaults() {
	if _, ok := vmc.mutation.CreatedAt(); !ok {
		v := virtualmachine.DefaultCreatedAt()
		vmc.mutation.SetCreatedAt(v)
	}
	if _, ok := vmc.mutation.UpdatedAt(); !ok {
		v := virtualmachine.DefaultUpdatedAt()
		vmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vmc.mutation.ID(); !ok {
		v := virtualmachine.DefaultID()
		vmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vmc *VirtualMachineCreate) check() error {
	if _, ok := vmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "VirtualMachine.created_at"`)}
	}
	if _, ok := vmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "VirtualMachine.updated_at"`)}
	}
	if _, ok := vmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "VirtualMachine.name"`)}
	}
	if v, ok := vmc.mutation.Name(); ok {
		if err := virtualmachine.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "VirtualMachine.name": %w`, err)}
		}
	}
	if _, ok := vmc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "VirtualMachine.owner_id"`)}
	}
	if _, ok := vmc.mutation.LocationID(); !ok {
		return &ValidationError{Name: "location_id", err: errors.New(`generated: missing required field "VirtualMachine.location_id"`)}
	}
	if v, ok := vmc.mutation.LocationID(); ok {
		if err := virtualmachine.LocationIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "location_id", err: fmt.Errorf(`generated: validator failed for field "VirtualMachine.location_id": %w`, err)}
		}
	}
	if _, ok := vmc.mutation.VMCPUConfigID(); !ok {
		return &ValidationError{Name: "vm_cpu_config_id", err: errors.New(`generated: missing required field "VirtualMachine.vm_cpu_config_id"`)}
	}
	if _, ok := vmc.mutation.VirtualMachineCPUConfigID(); !ok {
		return &ValidationError{Name: "virtual_machine_cpu_config", err: errors.New(`generated: missing required edge "VirtualMachine.virtual_machine_cpu_config"`)}
	}
	return nil
}

func (vmc *VirtualMachineCreate) sqlSave(ctx context.Context) (*VirtualMachine, error) {
	if err := vmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	vmc.mutation.id = &_node.ID
	vmc.mutation.done = true
	return _node, nil
}

func (vmc *VirtualMachineCreate) createSpec() (*VirtualMachine, *sqlgraph.CreateSpec) {
	var (
		_node = &VirtualMachine{config: vmc.config}
		_spec = sqlgraph.NewCreateSpec(virtualmachine.Table, sqlgraph.NewFieldSpec(virtualmachine.FieldID, field.TypeString))
	)
	if id, ok := vmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vmc.mutation.CreatedAt(); ok {
		_spec.SetField(virtualmachine.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vmc.mutation.UpdatedAt(); ok {
		_spec.SetField(virtualmachine.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vmc.mutation.Name(); ok {
		_spec.SetField(virtualmachine.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vmc.mutation.OwnerID(); ok {
		_spec.SetField(virtualmachine.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := vmc.mutation.LocationID(); ok {
		_spec.SetField(virtualmachine.FieldLocationID, field.TypeString, value)
		_node.LocationID = value
	}
	if value, ok := vmc.mutation.Userdata(); ok {
		_spec.SetField(virtualmachine.FieldUserdata, field.TypeString, value)
		_node.Userdata = value
	}
	if nodes := vmc.mutation.VirtualMachineCPUConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   virtualmachine.VirtualMachineCPUConfigTable,
			Columns: []string{virtualmachine.VirtualMachineCPUConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(virtualmachinecpuconfig.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VMCPUConfigID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VirtualMachineCreateBulk is the builder for creating many VirtualMachine entities in bulk.
type VirtualMachineCreateBulk struct {
	config
	err      error
	builders []*VirtualMachineCreate
}

// Save creates the VirtualMachine entities in the database.
func (vmcb *VirtualMachineCreateBulk) Save(ctx context.Context) ([]*VirtualMachine, error) {
	if vmcb.err != nil {
		return nil, vmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vmcb.builders))
	nodes := make([]*VirtualMachine, len(vmcb.builders))
	mutators := make([]Mutator, len(vmcb.builders))
	for i := range vmcb.builders {
		func(i int, root context.Context) {
			builder := vmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VirtualMachineMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vmcb *VirtualMachineCreateBulk) SaveX(ctx context.Context) []*VirtualMachine {
	v, err := vmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmcb *VirtualMachineCreateBulk) Exec(ctx context.Context) error {
	_, err := vmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmcb *VirtualMachineCreateBulk) ExecX(ctx context.Context) {
	if err := vmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
