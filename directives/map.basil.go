// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// MapInterpreter is the basil interpreter for the Map block
type MapInterpreter struct {
	s schema.Schema
}

func (i MapInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"eval_stage": "parse"},
			},
			Name: "Map",
			Properties: map[string]schema.Schema{
				"additional_properties": &schema.Reference{
					Ref: "http://basil.schema/github.com/opsidian/basil/basil/schema/Schema",
				},
				"annotations": &schema.Map{
					AdditionalProperties: &schema.String{},
				},
				"const": &schema.Map{
					AdditionalProperties: &schema.Untyped{},
				},
				"default": &schema.Map{
					AdditionalProperties: &schema.Untyped{},
				},
				"deprecated":  &schema.Boolean{},
				"description": &schema.String{},
				"enum": &schema.Array{
					Items: &schema.Map{
						AdditionalProperties: &schema.Untyped{},
					},
				},
				"examples": &schema.Array{
					Items: &schema.Untyped{},
				},
				"pointer":    &schema.Boolean{},
				"read_only":  &schema.Boolean{},
				"title":      &schema.String{},
				"write_only": &schema.Boolean{},
			},
			PropertyNames: map[string]string{"additional_properties": "AdditionalProperties", "annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "pointer": "Pointer", "read_only": "ReadOnly", "title": "Title", "write_only": "WriteOnly"},
			Required:      []string{"additional_properties"},
		}
	}
	return i.s
}

// Create creates a new Map block
func (i MapInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Map{}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i MapInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i MapInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Map
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MapInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "annotations":
		return b.(*Map).Annotations
	case "const":
		return b.(*Map).Const
	case "default":
		return b.(*Map).Default
	case "deprecated":
		return b.(*Map).Deprecated
	case "description":
		return b.(*Map).Description
	case "enum":
		return b.(*Map).Enum
	case "examples":
		return b.(*Map).Examples
	case "pointer":
		return b.(*Map).Pointer
	case "read_only":
		return b.(*Map).ReadOnly
	case "title":
		return b.(*Map).Title
	case "write_only":
		return b.(*Map).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in Map", name))
	}
}

func (i MapInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Map)
	switch name {
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = value.(map[string]interface{})
	case "default":
		b.Default = value.(map[string]interface{})
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([]map[string]interface{}, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.(map[string]interface{})
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "pointer":
		b.Pointer = value.(bool)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i MapInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Map)
	switch name {
	case "additional_properties":
		b.AdditionalProperties = value.(schema.Schema)
	}
	return nil
}
