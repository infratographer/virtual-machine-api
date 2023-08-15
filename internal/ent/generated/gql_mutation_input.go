
        

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
	"math"
	"strings"
	"sync"
	"time"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"
	"github.com/99designs/gqlgen/graphql"


)



import (
			"go.infratographer.com/x/gidx"
        "go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
)
    // CreateVirtualMachineInput represents a mutation input for creating virtualmachines.
    type CreateVirtualMachineInput struct {
            Name string
            OwnerID gidx.PrefixedID
            LocationID gidx.PrefixedID
            Userdata string
    }

    // Mutate applies the CreateVirtualMachineInput on the VirtualMachineMutation builder.
    func (i *CreateVirtualMachineInput) Mutate(m *VirtualMachineMutation) {
                m.SetName(i.Name)
                m.SetOwnerID(i.OwnerID)
                m.SetLocationID(i.LocationID)
                m.SetUserdata(i.Userdata)
    }
    // SetInput applies the change-set in the CreateVirtualMachineInput on the VirtualMachineCreate builder.
    func(c *VirtualMachineCreate) SetInput(i CreateVirtualMachineInput) *VirtualMachineCreate {
        i.Mutate(c.Mutation())
        return c
    }
    // UpdateVirtualMachineInput represents a mutation input for updating virtualmachines.
    type UpdateVirtualMachineInput struct {
            Name *string
            Userdata *string
    }

    // Mutate applies the UpdateVirtualMachineInput on the VirtualMachineMutation builder.
    func (i *UpdateVirtualMachineInput) Mutate(m *VirtualMachineMutation) {
                if v := i.Name; v != nil {
                    m.SetName(*v)
                }
                if v := i.Userdata; v != nil {
                    m.SetUserdata(*v)
                }
    }
    // SetInput applies the change-set in the UpdateVirtualMachineInput on the VirtualMachineUpdate builder.
    func(c *VirtualMachineUpdate) SetInput(i UpdateVirtualMachineInput) *VirtualMachineUpdate {
        i.Mutate(c.Mutation())
        return c
    }
    // SetInput applies the change-set in the UpdateVirtualMachineInput on the VirtualMachineUpdateOne builder.
    func(c *VirtualMachineUpdateOne) SetInput(i UpdateVirtualMachineInput) *VirtualMachineUpdateOne {
        i.Mutate(c.Mutation())
        return c
    }
