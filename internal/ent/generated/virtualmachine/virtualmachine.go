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

package virtualmachine

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the virtualmachine type in the database.
	Label = "virtual_machine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldUserdata holds the string denoting the userdata field in the database.
	FieldUserdata = "userdata"
	// FieldCores holds the string denoting the cores field in the database.
	FieldCores = "cores"
	// FieldSockets holds the string denoting the sockets field in the database.
	FieldSockets = "sockets"
	// Table holds the table name of the virtualmachine in the database.
	Table = "virtual_machines"
)

// Columns holds all SQL columns for virtualmachine fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldOwnerID,
	FieldLocationID,
	FieldUserdata,
	FieldCores,
	FieldSockets,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	LocationIDValidator func(string) error
	// DefaultCores holds the default value on creation for the "cores" field.
	DefaultCores int
	// CoresValidator is a validator for the "cores" field. It is called by the builders before save.
	CoresValidator func(int) error
	// DefaultSockets holds the default value on creation for the "sockets" field.
	DefaultSockets int
	// SocketsValidator is a validator for the "sockets" field. It is called by the builders before save.
	SocketsValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the VirtualMachine queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByUserdata orders the results by the userdata field.
func ByUserdata(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserdata, opts...).ToFunc()
}

// ByCores orders the results by the cores field.
func ByCores(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCores, opts...).ToFunc()
}

// BySockets orders the results by the sockets field.
func BySockets(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSockets, opts...).ToFunc()
}
