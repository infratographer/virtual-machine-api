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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachinecpuconfig"
	"go.infratographer.com/x/gidx"
)

// Represents a virtual machine on the graph.
type VirtualMachine struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the VirtualMachine.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the Virtual Machine.
	Name string `json:"name,omitempty"`
	// The ID for the owner of this Virtual Machine.
	OwnerID gidx.PrefixedID `json:"owner_id,omitempty"`
	// The ID for the location of this virtual machine.
	LocationID gidx.PrefixedID `json:"location_id,omitempty"`
	// The userdata for this virtual machine.
	Userdata string `json:"userdata,omitempty"`
	// The ID for the virtual machine cpu config.
	VMCPUConfigID gidx.PrefixedID `json:"vm_cpu_config_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VirtualMachineQuery when eager-loading is set.
	Edges        VirtualMachineEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VirtualMachineEdges holds the relations/edges for other nodes in the graph.
type VirtualMachineEdges struct {
	// The virtual machine cpu config for the virtual machine.
	VirtualMachineCPUConfig *VirtualMachineCPUConfig `json:"virtual_machine_cpu_config,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// VirtualMachineCPUConfigOrErr returns the VirtualMachineCPUConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VirtualMachineEdges) VirtualMachineCPUConfigOrErr() (*VirtualMachineCPUConfig, error) {
	if e.loadedTypes[0] {
		if e.VirtualMachineCPUConfig == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: virtualmachinecpuconfig.Label}
		}
		return e.VirtualMachineCPUConfig, nil
	}
	return nil, &NotLoadedError{edge: "virtual_machine_cpu_config"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VirtualMachine) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case virtualmachine.FieldID, virtualmachine.FieldOwnerID, virtualmachine.FieldLocationID, virtualmachine.FieldVMCPUConfigID:
			values[i] = new(gidx.PrefixedID)
		case virtualmachine.FieldName, virtualmachine.FieldUserdata:
			values[i] = new(sql.NullString)
		case virtualmachine.FieldCreatedAt, virtualmachine.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VirtualMachine fields.
func (vm *VirtualMachine) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case virtualmachine.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vm.ID = *value
			}
		case virtualmachine.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vm.CreatedAt = value.Time
			}
		case virtualmachine.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vm.UpdatedAt = value.Time
			}
		case virtualmachine.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				vm.Name = value.String
			}
		case virtualmachine.FieldOwnerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				vm.OwnerID = *value
			}
		case virtualmachine.FieldLocationID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value != nil {
				vm.LocationID = *value
			}
		case virtualmachine.FieldUserdata:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userdata", values[i])
			} else if value.Valid {
				vm.Userdata = value.String
			}
		case virtualmachine.FieldVMCPUConfigID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field vm_cpu_config_id", values[i])
			} else if value != nil {
				vm.VMCPUConfigID = *value
			}
		default:
			vm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VirtualMachine.
// This includes values selected through modifiers, order, etc.
func (vm *VirtualMachine) Value(name string) (ent.Value, error) {
	return vm.selectValues.Get(name)
}

// QueryVirtualMachineCPUConfig queries the "virtual_machine_cpu_config" edge of the VirtualMachine entity.
func (vm *VirtualMachine) QueryVirtualMachineCPUConfig() *VirtualMachineCPUConfigQuery {
	return NewVirtualMachineClient(vm.config).QueryVirtualMachineCPUConfig(vm)
}

// Update returns a builder for updating this VirtualMachine.
// Note that you need to call VirtualMachine.Unwrap() before calling this method if this VirtualMachine
// was returned from a transaction, and the transaction was committed or rolled back.
func (vm *VirtualMachine) Update() *VirtualMachineUpdateOne {
	return NewVirtualMachineClient(vm.config).UpdateOne(vm)
}

// Unwrap unwraps the VirtualMachine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vm *VirtualMachine) Unwrap() *VirtualMachine {
	_tx, ok := vm.config.driver.(*txDriver)
	if !ok {
		panic("generated: VirtualMachine is not a transactional entity")
	}
	vm.config.driver = _tx.drv
	return vm
}

// String implements the fmt.Stringer.
func (vm *VirtualMachine) String() string {
	var builder strings.Builder
	builder.WriteString("VirtualMachine(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(vm.Name)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", vm.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", vm.LocationID))
	builder.WriteString(", ")
	builder.WriteString("userdata=")
	builder.WriteString(vm.Userdata)
	builder.WriteString(", ")
	builder.WriteString("vm_cpu_config_id=")
	builder.WriteString(fmt.Sprintf("%v", vm.VMCPUConfigID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (vm VirtualMachine) IsEntity() {}

// VirtualMachines is a parsable slice of VirtualMachine.
type VirtualMachines []*VirtualMachine
