// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type BlockWithReferenceInterpreter struct{}

// Create creates a new BlockWithReference block
func (i BlockWithReferenceInterpreter) Create(ctx *basil.EvalContext, node basil.BlockNode) basil.Block {
	return &BlockWithReference{
		IDField: node.ID(),
	}
}

// Params returns with the list of valid parameters
func (i BlockWithReferenceInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockWithReferenceInterpreter) HasForeignID() bool {
	return true
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockWithReferenceInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithReferenceInterpreter) ParseContext(parentCtx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithReference
	if b, ok := basil.Block(nilBlock).(basil.ParseContextAware); ok {
		return b.ParseContext(parentCtx)
	}

	return parentCtx
}

func (i BlockWithReferenceInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockWithReference).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithReference", name))
	}
}

func (i BlockWithReferenceInterpreter) SetParam(ctx *basil.EvalContext, b basil.Block, name basil.ID, node basil.ParameterNode) parsley.Error {
	return nil
}

func (i BlockWithReferenceInterpreter) SetBlock(ctx *basil.EvalContext, b basil.Block, name basil.ID, value interface{}) parsley.Error {
	return nil
}