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

package pubsubhooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/hook"
	"go.infratographer.com/virtual-machine-api/internal/ent/schema"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
	"golang.org/x/exp/slices"
)

func VirtualMachineHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.VirtualMachineFunc(func(ctx context.Context, m *generated.VirtualMachineMutation) (ent.Value, error) {
					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					// queueName := ""
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_name := ""
					name, ok := m.Name()

					if ok {
						cv_name = fmt.Sprintf("%s", fmt.Sprint(name))
						pv_name := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldName(ctx)
							if err != nil {
								pv_name = "<unknown>"
							} else {
								pv_name = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "name",
							PreviousValue: pv_name,
							CurrentValue:  cv_name,
						})
					}

					cv_owner_id := ""
					owner_id, ok := m.OwnerID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						owner_id, err = m.OldOwnerID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, owner_id)

					if ok {
						cv_owner_id = fmt.Sprintf("%s", fmt.Sprint(owner_id))
						pv_owner_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldOwnerID(ctx)
							if err != nil {
								pv_owner_id = "<unknown>"
							} else {
								pv_owner_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "owner_id",
							PreviousValue: pv_owner_id,
							CurrentValue:  cv_owner_id,
						})
					}

					cv_location_id := ""
					location_id, ok := m.LocationID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						location_id, err = m.OldLocationID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, location_id)

					if ok {
						cv_location_id = fmt.Sprintf("%s", fmt.Sprint(location_id))
						pv_location_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldLocationID(ctx)
							if err != nil {
								pv_location_id = "<unknown>"
							} else {
								pv_location_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "location_id",
							PreviousValue: pv_location_id,
							CurrentValue:  cv_location_id,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					vm_lookup := getLocation(ctx, objID, additionalSubjects)
					if vm_lookup != "" {
						vm, err := m.Client().VirtualMachine.Get(ctx, vm_lookup)
						if err != nil {
							return nil, fmt.Errorf("unable to lookup location %s", vm_lookup)
						}

						if !slices.Contains(additionalSubjects, vm.LocationID) {
							additionalSubjects = append(additionalSubjects, vm.LocationID)
							msg.AdditionalSubjectIDs = additionalSubjects
						}
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.VirtualMachineFunc(func(ctx context.Context, m *generated.VirtualMachineMutation) (ent.Value, error) {
					// queueName := ""
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().VirtualMachine.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for pubsub, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.OwnerID)

					additionalSubjects = append(additionalSubjects, dbObj.LocationID)

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					vm_lookup := getLocation(ctx, objID, additionalSubjects)
					if vm_lookup != "" {
						vm, err := m.Client().VirtualMachine.Get(ctx, vm_lookup)
						if err != nil {
							return nil, fmt.Errorf("unable to lookup location %s", vm_lookup)
						}

						if !slices.Contains(additionalSubjects, vm.LocationID) {
							additionalSubjects = append(additionalSubjects, vm.LocationID)
							msg.AdditionalSubjectIDs = additionalSubjects
						}
					}

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

func PubsubHooks(c *generated.Client) {
	c.VirtualMachine.Use(VirtualMachineHooks()...)

}

func eventType(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return "create"
	case ent.OpUpdate, ent.OpUpdateOne:
		return "update"
	case ent.OpDelete, ent.OpDeleteOne:
		return "delete"
	default:
		return "unknown"
	}
}

func eventSubject(objID gidx.PrefixedID) string {
	switch objID.Prefix() {
	case schema.VirtualMachinePrefix:
		return "virtual-machine"
	/*
	 * This is intentionall commented until the package is generalized in infratographer/x pkg
	 */
	// case schema.PortPrefix:
	//         return "load-balancer-port"
	// case schema.OriginPrefix:
	//         return "load-balancer-origin"
	// case schema.PoolPrefix:
	//         return "load-balancer-pool"
	default:
		return "unknown"
	}
}

func getLocation(ctx context.Context, id gidx.PrefixedID, addID []gidx.PrefixedID) gidx.PrefixedID {
	if id.Prefix() == schema.VirtualMachinePrefix {
		return id
	}

	for _, id := range addID {
		if id.Prefix() == schema.VirtualMachinePrefix {
			return id
		}
	}

	return ""
}
