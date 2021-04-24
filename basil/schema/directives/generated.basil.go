// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// GeneratedInterpreter is the basil interpreter for the Generated block
type GeneratedInterpreter struct {
	s schema.Schema
}

func (i GeneratedInterpreter) Schema() schema.Schema {
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

// Create creates a new Generated block
func (i GeneratedInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Generated{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i GeneratedInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i GeneratedInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Generated
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i GeneratedInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Generated).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Generated", name))
	}
}

func (i GeneratedInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i GeneratedInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
