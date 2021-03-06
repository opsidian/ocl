// Code generated by Basil. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// BlockGeneratorResultInterpreter is the basil interpreter for the BlockGeneratorResult block
type BlockGeneratorResultInterpreter struct {
	s schema.Schema
}

func (i BlockGeneratorResultInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BlockGeneratorResult",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"value": &schema.Untyped{},
			},
		}
	}
	return i.s
}

// Create creates a new BlockGeneratorResult block
func (i BlockGeneratorResultInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BlockGeneratorResult{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockGeneratorResultInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockGeneratorResultInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockGeneratorResult
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockGeneratorResultInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockGeneratorResult).id
	case "value":
		return b.(*BlockGeneratorResult).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockGeneratorResult", name))
	}
}

func (i BlockGeneratorResultInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*BlockGeneratorResult)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i BlockGeneratorResultInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
