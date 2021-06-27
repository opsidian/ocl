// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// NameInterpreter is the basil interpreter for the Name block
type NameInterpreter struct {
	s schema.Schema
}

func (i NameInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Name",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"value": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
			},
			PropertyNames: map[string]string{"value": "Value"},
			Required:      []string{"value"},
		}
	}
	return i.s
}

// Create creates a new Name block
func (i NameInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Name{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i NameInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i NameInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Name
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i NameInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Name).id
	case "value":
		return b.(*Name).Value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Name", name))
	}
}

func (i NameInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Name)
	switch name {
	case "value":
		b.Value = value.(string)
	}
	return nil
}

func (i NameInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}