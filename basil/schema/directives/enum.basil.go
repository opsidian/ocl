// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// EnumInterpreter is the basil interpreter for the Enum block
type EnumInterpreter struct {
	s schema.Schema
}

func (i EnumInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Enum",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"values": &schema.Array{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
					Items: &schema.Untyped{},
				},
			},
			Required: []string{"values"},
		}
	}
	return i.s
}

// Create creates a new Enum block
func (i EnumInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Enum{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i EnumInterpreter) ValueParamName() basil.ID {
	return "values"
}

// ParseContext returns with the parse context for the block
func (i EnumInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Enum
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i EnumInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Enum).id
	case "values":
		return b.(*Enum).values
	default:
		panic(fmt.Errorf("unexpected parameter %q in Enum", name))
	}
}

func (i EnumInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Enum)
	switch name {
	case "values":
		b.values = value.([]interface{})
	}
	return nil
}

func (i EnumInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
