// Code generated by Basil. DO NOT EDIT.
package blocks

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type FailInterpreter struct{}

// Create creates a new Fail block
func (i FailInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Fail{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i FailInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"msg": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i FailInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i FailInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i FailInterpreter) ValueParamName() basil.ID {
	return "msg"
}

// ParseContext returns with the parse context for the block
func (i FailInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Fail
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i FailInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Fail).id
	case "msg":
		return b.(*Fail).msg
	default:
		panic(fmt.Errorf("unexpected parameter %q in Fail", name))
	}
}

func (i FailInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Fail)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "msg":
		b.msg, err = variable.StringValue(value)
	}
	return err
}

func (i FailInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i FailInterpreter) EvalStage() basil.EvalStage {
	var nilBlock *Fail
	if b, ok := basil.Block(nilBlock).(basil.EvalStageAware); ok {
		return b.EvalStage()
	}

	return basil.EvalStageUndefined
}
