// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "code", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sequence", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString, Size: 20, Default: ""},
		{Name: "path", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "properties", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"disabled", "enabled"}, Default: "disabled"},
		{Name: "parent_path", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "parent_id", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[10]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menu_code",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[1]},
			},
			{
				Name:    "menu_name",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[2]},
			},
			{
				Name:    "menu_type",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[5]},
			},
			{
				Name:    "menu_status",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[8]},
			},
			{
				Name:    "menu_parent_id",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[10]},
			},
			{
				Name:    "menu_parent_path",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[9]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysMenusTable,
	}
)

func init() {
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
}
