// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// VirtMsColumns holds the columns for the "virt_ms" table.
	VirtMsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "hostname", Type: field.TypeString, Size: 2147483647},
	}
	// VirtMsTable holds the schema information for the "virt_ms" table.
	VirtMsTable = &schema.Table{
		Name:       "virt_ms",
		Columns:    VirtMsColumns,
		PrimaryKey: []*schema.Column{VirtMsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "virtm_created_at",
				Unique:  false,
				Columns: []*schema.Column{VirtMsColumns[1]},
			},
			{
				Name:    "virtm_updated_at",
				Unique:  false,
				Columns: []*schema.Column{VirtMsColumns[2]},
			},
			{
				Name:    "virtm_hostname",
				Unique:  false,
				Columns: []*schema.Column{VirtMsColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		VirtMsTable,
	}
)

func init() {
}
