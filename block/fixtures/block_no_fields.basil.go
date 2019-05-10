// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type BlockNoFieldsInterpreter struct{}

// Create creates a new BlockNoFields block
func (i BlockNoFieldsInterpreter) Create(ctx *basil.EvalContext, node basil.BlockNode) basil.Block {
	return &BlockNoFields{
		IDField: node.ID(),
	}
}

// Params returns with the list of valid parameters
func (i BlockNoFieldsInterpreter) Params() map[basil.ID]string {
	return nil
}

// RequiredParams returns with the list of required parameters
func (i BlockNoFieldsInterpreter) RequiredParams() map[basil.ID]bool {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockNoFieldsInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockNoFieldsInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockNoFieldsInterpreter) ParseContext(parentCtx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockNoFields
	if b, ok := basil.Block(nilBlock).(basil.ParseContextAware); ok {
		return b.ParseContext(parentCtx)
	}

	return parentCtx
}

func (i BlockNoFieldsInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockNoFields", name))
	}
}

func (i BlockNoFieldsInterpreter) SetParam(ctx *basil.EvalContext, b basil.Block, name basil.ID, node basil.BlockParamNode) parsley.Error {
	switch name {
	}

	return nil
}

func (i BlockNoFieldsInterpreter) SetBlock(ctx *basil.EvalContext, b basil.Block, name basil.ID, value interface{}) parsley.Error {
	return nil
}
