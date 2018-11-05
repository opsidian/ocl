// Code generated by Basil. DO NOT EDIT.
package test

import (
	"fmt"
	"strings"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/parsley/parsley"
)

// NewTestBlockFactory creates a new TestBlock block factory
func NewTestBlockFactory(
	typeNode parsley.Node,
	idNode parsley.Node,
	paramNodes map[basil.ID]parsley.Node,
	blockNodes []parsley.Node,
) (basil.BlockFactory, parsley.Error) {
	return &TestBlockFactory{
		typeNode:   typeNode,
		idNode:     idNode,
		paramNodes: paramNodes,
		blockNodes: blockNodes,
	}, nil
}

// TestBlockFactory will create and evaluate a TestBlock block
type TestBlockFactory struct {
	typeNode    parsley.Node
	idNode      parsley.Node
	paramNodes  map[basil.ID]parsley.Node
	blockNodes  []parsley.Node
	shortFormat bool
}

// CreateBlock creates a new TestBlock block
func (f *TestBlockFactory) CreateBlock(parentCtx interface{}) (basil.Block, interface{}, parsley.Error) {
	var err parsley.Error

	block := &TestBlock{}

	if block.IDField, err = util.NodeIdentifierValue(f.idNode, parentCtx); err != nil {
		return nil, nil, err
	}

	ctx := block.Context(parentCtx)

	if valueNode, ok := f.paramNodes["_value"]; ok {
		f.shortFormat = true

		if block.Value, err = util.NodeAnyValue(valueNode, ctx); err != nil {
			return nil, nil, err
		}
	}

	if len(f.blockNodes) > 0 {
		var childBlockFactory interface{}
		for _, childBlock := range f.blockNodes {
			if childBlockFactory, err = childBlock.Value(ctx); err != nil {
				return nil, nil, err
			}
			switch b := childBlockFactory.(type) {
			case *TestBlockFactory:
				block.BlockFactories = append(block.BlockFactories, b)
			default:
				panic(fmt.Sprintf("block type %T is not supported in TestBlock, please open a bug ticket", childBlockFactory))
			}

		}
	}

	return block, ctx, nil
}

// EvalBlock evaluates all fields belonging to the given stage on a TestBlock block
func (f *TestBlockFactory) EvalBlock(ctx interface{}, stage string, res basil.Block) parsley.Error {
	var err parsley.Error

	if preInterpreter, ok := res.(basil.BlockPreInterpreter); ok {
		if err = preInterpreter.PreEval(ctx, stage); err != nil {
			return err
		}
	}

	block, ok := res.(*TestBlock)
	if !ok {
		panic("result must be a type of *TestBlock")
	}

	validParamNames := map[basil.ID]struct{}{
		"value":               struct{}{},
		"field_string":        struct{}{},
		"field_int":           struct{}{},
		"field_float":         struct{}{},
		"field_bool":          struct{}{},
		"field_array":         struct{}{},
		"field_map":           struct{}{},
		"field_time_duration": struct{}{},
		"custom_field":        struct{}{},
		"block_factories":     struct{}{},
	}

	for paramName, paramNode := range f.paramNodes {
		if !strings.HasPrefix(string(paramName), "_") {
			if _, valid := validParamNames[paramName]; !valid {
				return parsley.NewError(paramNode.Pos(), fmt.Errorf("%q parameter does not exist", paramName))
			}
		}
	}

	if !f.shortFormat {
		switch stage {
		case "default":
			if valueNode, ok := f.paramNodes["value"]; ok {
				if block.Value, err = util.NodeAnyValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_string"]; ok {
				if block.FieldString, err = util.NodeStringValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_int"]; ok {
				if block.FieldInt, err = util.NodeIntegerValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_float"]; ok {
				if block.FieldFloat, err = util.NodeFloatValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_bool"]; ok {
				if block.FieldBool, err = util.NodeBoolValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_array"]; ok {
				if block.FieldArray, err = util.NodeArrayValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_map"]; ok {
				if block.FieldMap, err = util.NodeMapValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["field_time_duration"]; ok {
				if block.FieldTimeDuration, err = util.NodeTimeDurationValue(valueNode, ctx); err != nil {
					return err
				}
			}

			if valueNode, ok := f.paramNodes["custom_field"]; ok {
				if block.FieldCustomName, err = util.NodeStringValue(valueNode, ctx); err != nil {
					return err
				}
			}
		default:
			panic(fmt.Sprintf("unknown stage: %s", stage))
		}

		switch stage {
		case "default":
			var childBlock basil.Block
			var childBlockCtx interface{}
			for _, childBlockFactory := range block.BlockFactories {
				if childBlock, childBlockCtx, err = childBlockFactory.CreateBlock(ctx); err != nil {
					return err
				}

				if err = childBlockFactory.EvalBlock(childBlockCtx, stage, childBlock); err != nil {
					return err
				}

				switch b := childBlock.(type) {
				case *TestBlock:
					block.Blocks = append(block.Blocks, b)
				default:
					panic(fmt.Sprintf("block type %T is not supported in BlockFactories, please open a bug ticket", childBlock))
				}

			}
		default:
			panic(fmt.Sprintf("unknown stage: %s", stage))
		}
	}

	if postInterpreter, ok := res.(basil.BlockPostInterpreter); ok {
		if err = postInterpreter.PostEval(ctx, stage); err != nil {
			return err
		}
	}

	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (f *TestBlockFactory) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (f *TestBlockFactory) HasShortFormat() bool {
	return true
}