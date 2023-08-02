// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testclient

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"go.infratographer.com/x/gidx"
)

// An object with an ID.
// Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
type Node interface {
	IsNode()
	// The id of the object.
	GetID() gidx.PrefixedID
}

type Entity interface {
	IsEntity()
}

// Create a new virtual machine.
type CreateVirtMInput struct {
	// The name of the VirtM.
	Hostname string `json:"hostname"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

// Update an existing virtual machine.
type UpdateVirtMInput struct {
	// The name of the VirtM.
	Hostname *string `json:"hostname,omitempty"`
}

type VirtM struct {
	// The ID of the VirtM.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// The name of the VirtM.
	Hostname string `json:"hostname"`
}

func (VirtM) IsNode() {}

// The id of the object.
func (this VirtM) GetID() gidx.PrefixedID { return this.ID }

func (VirtM) IsEntity() {}

// A connection to a list of items.
type VirtMConnection struct {
	// A list of edges.
	Edges []*VirtMEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type VirtMEdge struct {
	// The item at the end of the edge.
	Node *VirtM `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for VirtM connections
type VirtMOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order VirtMs.
	Field VirtMOrderField `json:"field"`
}

// VirtMWhereInput is used for filtering VirtM objects.
// Input was generated by ent.
type VirtMWhereInput struct {
	Not *VirtMWhereInput   `json:"not,omitempty"`
	And []*VirtMWhereInput `json:"and,omitempty"`
	Or  []*VirtMWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// hostname field predicates
	Hostname             *string  `json:"hostname,omitempty"`
	HostnameNeq          *string  `json:"hostnameNEQ,omitempty"`
	HostnameIn           []string `json:"hostnameIn,omitempty"`
	HostnameNotIn        []string `json:"hostnameNotIn,omitempty"`
	HostnameGt           *string  `json:"hostnameGT,omitempty"`
	HostnameGte          *string  `json:"hostnameGTE,omitempty"`
	HostnameLt           *string  `json:"hostnameLT,omitempty"`
	HostnameLte          *string  `json:"hostnameLTE,omitempty"`
	HostnameContains     *string  `json:"hostnameContains,omitempty"`
	HostnameHasPrefix    *string  `json:"hostnameHasPrefix,omitempty"`
	HostnameHasSuffix    *string  `json:"hostnameHasSuffix,omitempty"`
	HostnameEqualFold    *string  `json:"hostnameEqualFold,omitempty"`
	HostnameContainsFold *string  `json:"hostnameContainsFold,omitempty"`
}

type Service struct {
	Sdl *string `json:"sdl,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which VirtM connections can be ordered.
type VirtMOrderField string

const (
	VirtMOrderFieldID        VirtMOrderField = "ID"
	VirtMOrderFieldCreatedAt VirtMOrderField = "CREATED_AT"
	VirtMOrderFieldUpdatedAt VirtMOrderField = "UPDATED_AT"
	VirtMOrderFieldHostname  VirtMOrderField = "HOSTNAME"
)

var AllVirtMOrderField = []VirtMOrderField{
	VirtMOrderFieldID,
	VirtMOrderFieldCreatedAt,
	VirtMOrderFieldUpdatedAt,
	VirtMOrderFieldHostname,
}

func (e VirtMOrderField) IsValid() bool {
	switch e {
	case VirtMOrderFieldID, VirtMOrderFieldCreatedAt, VirtMOrderFieldUpdatedAt, VirtMOrderFieldHostname:
		return true
	}
	return false
}

func (e VirtMOrderField) String() string {
	return string(e)
}

func (e *VirtMOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VirtMOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VirtMOrderField", str)
	}
	return nil
}

func (e VirtMOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
