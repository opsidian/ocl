// Code generated by Basil. DO NOT EDIT.

package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// BlockWithManyBlockInterpreter is the basil interpreter for the BlockWithManyBlock block
type BlockWithManyBlockInterpreter struct {
	s schema.Schema
}

func (i BlockWithManyBlockInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Properties: map[string]schema.Schema{
				"block": &schema.Array{
					Items: &schema.Reference{
						Metadata: schema.Metadata{
							Pointer: true,
						},
						Ref: "http://basil.schema/github.com/opsidian/basil/test/fixtures/Block",
					},
				},
				"id_field": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
			StructProperties: map[string]string{"block": "Block", "id_field": "IDField"},
		}
	}
	return i.s
}

// Create creates a new BlockWithManyBlock block
func (i BlockWithManyBlockInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BlockWithManyBlock{
		IDField: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockWithManyBlockInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithManyBlockInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithManyBlock
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithManyBlockInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockWithManyBlock).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithManyBlock", name))
	}
}

func (i BlockWithManyBlockInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i BlockWithManyBlockInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*BlockWithManyBlock)
	switch name {
	case "block":
		b.Block = append(b.Block, value.(*Block))
	}
	return nil
}
