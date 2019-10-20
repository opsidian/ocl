// Code generated by Basil. DO NOT EDIT.
package directive

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type WhenInterpreter struct{}

// Create creates a new When block
func (i WhenInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &When{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i WhenInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"cond": {
			Type:       "bool",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i WhenInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i WhenInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i WhenInterpreter) ValueParamName() basil.ID {
	return "cond"
}

// ParseContext returns with the parse context for the block
func (i WhenInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *When
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i WhenInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*When).id
	case "cond":
		return b.(*When).cond
	default:
		panic(fmt.Errorf("unexpected parameter %q in When", name))
	}
}

func (i WhenInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*When)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "cond":
		b.cond, err = variable.BoolValue(value)
	}
	return err
}

func (i WhenInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
