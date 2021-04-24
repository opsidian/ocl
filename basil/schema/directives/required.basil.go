// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// RequiredInterpreter is the basil interpreter for the Required block
type RequiredInterpreter struct {
	s schema.Schema
}

func (i RequiredInterpreter) Schema() schema.Schema {
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
			},
		}
	}
	return i.s
}

// Create creates a new Required block
func (i RequiredInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Required{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i RequiredInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i RequiredInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Required
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i RequiredInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Required).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Required", name))
	}
}

func (i RequiredInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i RequiredInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
