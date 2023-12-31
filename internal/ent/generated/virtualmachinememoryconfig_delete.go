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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/predicate"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinememoryconfig"
)

// VirtualMachineMemoryConfigDelete is the builder for deleting a VirtualMachineMemoryConfig entity.
type VirtualMachineMemoryConfigDelete struct {
	config
	hooks    []Hook
	mutation *VirtualMachineMemoryConfigMutation
}

// Where appends a list predicates to the VirtualMachineMemoryConfigDelete builder.
func (vmmcd *VirtualMachineMemoryConfigDelete) Where(ps ...predicate.VirtualMachineMemoryConfig) *VirtualMachineMemoryConfigDelete {
	vmmcd.mutation.Where(ps...)
	return vmmcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vmmcd *VirtualMachineMemoryConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vmmcd.sqlExec, vmmcd.mutation, vmmcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vmmcd *VirtualMachineMemoryConfigDelete) ExecX(ctx context.Context) int {
	n, err := vmmcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vmmcd *VirtualMachineMemoryConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(virtualmachinememoryconfig.Table, sqlgraph.NewFieldSpec(virtualmachinememoryconfig.FieldID, field.TypeString))
	if ps := vmmcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vmmcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vmmcd.mutation.done = true
	return affected, err
}

// VirtualMachineMemoryConfigDeleteOne is the builder for deleting a single VirtualMachineMemoryConfig entity.
type VirtualMachineMemoryConfigDeleteOne struct {
	vmmcd *VirtualMachineMemoryConfigDelete
}

// Where appends a list predicates to the VirtualMachineMemoryConfigDelete builder.
func (vmmcdo *VirtualMachineMemoryConfigDeleteOne) Where(ps ...predicate.VirtualMachineMemoryConfig) *VirtualMachineMemoryConfigDeleteOne {
	vmmcdo.vmmcd.mutation.Where(ps...)
	return vmmcdo
}

// Exec executes the deletion query.
func (vmmcdo *VirtualMachineMemoryConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := vmmcdo.vmmcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{virtualmachinememoryconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vmmcdo *VirtualMachineMemoryConfigDeleteOne) ExecX(ctx context.Context) {
	if err := vmmcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
