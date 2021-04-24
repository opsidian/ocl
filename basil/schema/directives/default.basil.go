// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// DefaultInterpreter is the basil interpreter for the Default block
type DefaultInterpreter struct {
	s schema.Schema
}

func (i DefaultInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"value": &schema.Untyped{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
			},
			Required: []string{"value"},
		}
	}
	return i.s
}

// Create creates a new Default block
func (i DefaultInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Default{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i DefaultInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i DefaultInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Default
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i DefaultInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Default).id
	case "value":
		return b.(*Default).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Default", name))
	}
}

func (i DefaultInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Default)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i DefaultInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
