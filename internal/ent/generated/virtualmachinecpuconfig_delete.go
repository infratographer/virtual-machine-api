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
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinecpuconfig"
)

// VirtualMachineCPUConfigDelete is the builder for deleting a VirtualMachineCPUConfig entity.
type VirtualMachineCPUConfigDelete struct {
	config
	hooks    []Hook
	mutation *VirtualMachineCPUConfigMutation
}

// Where appends a list predicates to the VirtualMachineCPUConfigDelete builder.
func (vmccd *VirtualMachineCPUConfigDelete) Where(ps ...predicate.VirtualMachineCPUConfig) *VirtualMachineCPUConfigDelete {
	vmccd.mutation.Where(ps...)
	return vmccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vmccd *VirtualMachineCPUConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vmccd.sqlExec, vmccd.mutation, vmccd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vmccd *VirtualMachineCPUConfigDelete) ExecX(ctx context.Context) int {
	n, err := vmccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vmccd *VirtualMachineCPUConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(virtualmachinecpuconfig.Table, sqlgraph.NewFieldSpec(virtualmachinecpuconfig.FieldID, field.TypeString))
	if ps := vmccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vmccd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vmccd.mutation.done = true
	return affected, err
}

// VirtualMachineCPUConfigDeleteOne is the builder for deleting a single VirtualMachineCPUConfig entity.
type VirtualMachineCPUConfigDeleteOne struct {
	vmccd *VirtualMachineCPUConfigDelete
}

// Where appends a list predicates to the VirtualMachineCPUConfigDelete builder.
func (vmccdo *VirtualMachineCPUConfigDeleteOne) Where(ps ...predicate.VirtualMachineCPUConfig) *VirtualMachineCPUConfigDeleteOne {
	vmccdo.vmccd.mutation.Where(ps...)
	return vmccdo
}

// Exec executes the deletion query.
func (vmccdo *VirtualMachineCPUConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := vmccdo.vmccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{virtualmachinecpuconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vmccdo *VirtualMachineCPUConfigDeleteOne) ExecX(ctx context.Context) {
	if err := vmccdo.Exec(ctx); err != nil {
		panic(err)
	}
}
