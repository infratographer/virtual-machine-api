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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinecpuconfig"
	"go.infratographer.com/x/gidx"
)

// VirtualMachineCPUConfigCreate is the builder for creating a VirtualMachineCPUConfig entity.
type VirtualMachineCPUConfigCreate struct {
	config
	mutation *VirtualMachineCPUConfigMutation
	hooks    []Hook
}

// SetCores sets the "cores" field.
func (vmccc *VirtualMachineCPUConfigCreate) SetCores(i int64) *VirtualMachineCPUConfigCreate {
	vmccc.mutation.SetCores(i)
	return vmccc
}

// SetSockets sets the "sockets" field.
func (vmccc *VirtualMachineCPUConfigCreate) SetSockets(i int64) *VirtualMachineCPUConfigCreate {
	vmccc.mutation.SetSockets(i)
	return vmccc
}

// SetID sets the "id" field.
func (vmccc *VirtualMachineCPUConfigCreate) SetID(gi gidx.PrefixedID) *VirtualMachineCPUConfigCreate {
	vmccc.mutation.SetID(gi)
	return vmccc
}

// SetVirtualMachineID sets the "virtual_machine" edge to the VirtualMachine entity by ID.
func (vmccc *VirtualMachineCPUConfigCreate) SetVirtualMachineID(id gidx.PrefixedID) *VirtualMachineCPUConfigCreate {
	vmccc.mutation.SetVirtualMachineID(id)
	return vmccc
}

// SetNillableVirtualMachineID sets the "virtual_machine" edge to the VirtualMachine entity by ID if the given value is not nil.
func (vmccc *VirtualMachineCPUConfigCreate) SetNillableVirtualMachineID(id *gidx.PrefixedID) *VirtualMachineCPUConfigCreate {
	if id != nil {
		vmccc = vmccc.SetVirtualMachineID(*id)
	}
	return vmccc
}

// SetVirtualMachine sets the "virtual_machine" edge to the VirtualMachine entity.
func (vmccc *VirtualMachineCPUConfigCreate) SetVirtualMachine(v *VirtualMachine) *VirtualMachineCPUConfigCreate {
	return vmccc.SetVirtualMachineID(v.ID)
}

// Mutation returns the VirtualMachineCPUConfigMutation object of the builder.
func (vmccc *VirtualMachineCPUConfigCreate) Mutation() *VirtualMachineCPUConfigMutation {
	return vmccc.mutation
}

// Save creates the VirtualMachineCPUConfig in the database.
func (vmccc *VirtualMachineCPUConfigCreate) Save(ctx context.Context) (*VirtualMachineCPUConfig, error) {
	return withHooks(ctx, vmccc.sqlSave, vmccc.mutation, vmccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vmccc *VirtualMachineCPUConfigCreate) SaveX(ctx context.Context) *VirtualMachineCPUConfig {
	v, err := vmccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmccc *VirtualMachineCPUConfigCreate) Exec(ctx context.Context) error {
	_, err := vmccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmccc *VirtualMachineCPUConfigCreate) ExecX(ctx context.Context) {
	if err := vmccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vmccc *VirtualMachineCPUConfigCreate) check() error {
	if _, ok := vmccc.mutation.Cores(); !ok {
		return &ValidationError{Name: "cores", err: errors.New(`generated: missing required field "VirtualMachineCPUConfig.cores"`)}
	}
	if v, ok := vmccc.mutation.Cores(); ok {
		if err := virtualmachinecpuconfig.CoresValidator(v); err != nil {
			return &ValidationError{Name: "cores", err: fmt.Errorf(`generated: validator failed for field "VirtualMachineCPUConfig.cores": %w`, err)}
		}
	}
	if _, ok := vmccc.mutation.Sockets(); !ok {
		return &ValidationError{Name: "sockets", err: errors.New(`generated: missing required field "VirtualMachineCPUConfig.sockets"`)}
	}
	if v, ok := vmccc.mutation.Sockets(); ok {
		if err := virtualmachinecpuconfig.SocketsValidator(v); err != nil {
			return &ValidationError{Name: "sockets", err: fmt.Errorf(`generated: validator failed for field "VirtualMachineCPUConfig.sockets": %w`, err)}
		}
	}
	return nil
}

func (vmccc *VirtualMachineCPUConfigCreate) sqlSave(ctx context.Context) (*VirtualMachineCPUConfig, error) {
	if err := vmccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vmccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vmccc.driver, _spec); err != nil {
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
	vmccc.mutation.id = &_node.ID
	vmccc.mutation.done = true
	return _node, nil
}

func (vmccc *VirtualMachineCPUConfigCreate) createSpec() (*VirtualMachineCPUConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &VirtualMachineCPUConfig{config: vmccc.config}
		_spec = sqlgraph.NewCreateSpec(virtualmachinecpuconfig.Table, sqlgraph.NewFieldSpec(virtualmachinecpuconfig.FieldID, field.TypeString))
	)
	if id, ok := vmccc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vmccc.mutation.Cores(); ok {
		_spec.SetField(virtualmachinecpuconfig.FieldCores, field.TypeInt64, value)
		_node.Cores = value
	}
	if value, ok := vmccc.mutation.Sockets(); ok {
		_spec.SetField(virtualmachinecpuconfig.FieldSockets, field.TypeInt64, value)
		_node.Sockets = value
	}
	if nodes := vmccc.mutation.VirtualMachineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   virtualmachinecpuconfig.VirtualMachineTable,
			Columns: []string{virtualmachinecpuconfig.VirtualMachineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(virtualmachine.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VirtualMachineCPUConfigCreateBulk is the builder for creating many VirtualMachineCPUConfig entities in bulk.
type VirtualMachineCPUConfigCreateBulk struct {
	config
	err      error
	builders []*VirtualMachineCPUConfigCreate
}

// Save creates the VirtualMachineCPUConfig entities in the database.
func (vmcccb *VirtualMachineCPUConfigCreateBulk) Save(ctx context.Context) ([]*VirtualMachineCPUConfig, error) {
	if vmcccb.err != nil {
		return nil, vmcccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vmcccb.builders))
	nodes := make([]*VirtualMachineCPUConfig, len(vmcccb.builders))
	mutators := make([]Mutator, len(vmcccb.builders))
	for i := range vmcccb.builders {
		func(i int, root context.Context) {
			builder := vmcccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VirtualMachineCPUConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, vmcccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vmcccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vmcccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vmcccb *VirtualMachineCPUConfigCreateBulk) SaveX(ctx context.Context) []*VirtualMachineCPUConfig {
	v, err := vmcccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmcccb *VirtualMachineCPUConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := vmcccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmcccb *VirtualMachineCPUConfigCreateBulk) ExecX(ctx context.Context) {
	if err := vmcccb.Exec(ctx); err != nil {
		panic(err)
	}
}
