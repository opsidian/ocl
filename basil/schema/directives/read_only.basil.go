// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// ReadOnlyInterpreter is the basil interpreter for the ReadOnly block
type ReadOnlyInterpreter struct {
	s schema.Schema
}

func (i ReadOnlyInterpreter) Schema() schema.Schema {
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

// Create creates a new ReadOnly block
func (i ReadOnlyInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &ReadOnly{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i ReadOnlyInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ReadOnlyInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *ReadOnly
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ReadOnlyInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*ReadOnly).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in ReadOnly", name))
	}
}

func (i ReadOnlyInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i ReadOnlyInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
