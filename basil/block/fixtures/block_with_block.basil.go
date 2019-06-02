// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
)

type BlockWithBlockInterpreter struct{}

// Create creates a new BlockWithBlock block
func (i BlockWithBlockInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BlockWithBlock{
		IDField: id,
	}
}

// Params returns with the list of valid parameters
func (i BlockWithBlockInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockWithBlockInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockWithBlockInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithBlockInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithBlock
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithBlockInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockWithBlock).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithBlock", name))
	}
}

func (i BlockWithBlockInterpreter) SetParam(b basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i BlockWithBlockInterpreter) SetBlock(b basil.Block, name basil.ID, value interface{}) error {
	switch name {
	case "block_simple":
		b.(*BlockWithBlock).BlockSimple = append(b.(*BlockWithBlock).BlockSimple, value.(*BlockSimple))
	}
	return nil
}
