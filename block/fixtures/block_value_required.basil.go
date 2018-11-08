// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/parsley/parsley"
)

type BlockValueRequiredInterpreter struct{}

func (i BlockValueRequiredInterpreter) StaticCheck(ctx interface{}, node basil.BlockNode) (string, parsley.Error) {
	validParamNames := map[basil.ID]struct{}{
		"value": struct{}{},
	}

	for paramName, paramNode := range node.ParamNodes() {
		if _, valid := validParamNames[paramName]; !valid {
			return "", parsley.NewError(paramNode.Pos(), fmt.Errorf("%q parameter does not exist", paramName))
		}
	}

	if paramNode, ok := node.ParamNodes()["value"]; ok {
		if err := util.CheckNodeType(paramNode, "interface{}"); err != nil {
			return "", err
		}
	}

	requiredParamNames := []basil.ID{
		"value",
	}

	for _, paramName := range requiredParamNames {
		if _, set := node.ParamNodes()[paramName]; !set {
			return "", parsley.NewError(node.Pos(), fmt.Errorf("%s parameter is required", paramName))
		}
	}

	return "*BlockValueRequired", nil
}

// CreateBlock creates a new BlockValueRequired block
func (i BlockValueRequiredInterpreter) Eval(parentCtx interface{}, node basil.BlockNode) (basil.Block, parsley.Error) {
	block := &BlockValueRequired{
		IDField: node.ID(),
	}

	ctx := block.Context(parentCtx)

	if err := i.EvalBlock(ctx, node, "default", block); err != nil {
		return nil, err
	}

	return block, nil
}

// EvalBlock evaluates all fields belonging to the given stage on a BlockValueRequired block
func (i BlockValueRequiredInterpreter) EvalBlock(ctx interface{}, node basil.BlockNode, stage string, res basil.Block) parsley.Error {
	var err parsley.Error

	if preInterpreter, ok := res.(basil.BlockPreInterpreter); ok {
		if err = preInterpreter.PreEval(ctx, stage); err != nil {
			return err
		}
	}

	block, ok := res.(*BlockValueRequired)
	if !ok {
		panic("result must be a type of *BlockValueRequired")
	}

	switch stage {
	case "default":
		if valueNode, ok := node.ParamNodes()["value"]; ok {
			if block.Value, err = util.NodeAnyValue(valueNode, ctx); err != nil {
				return err
			}
		}
	default:
		panic(fmt.Sprintf("unknown stage: %s", stage))
	}

	switch stage {
	case "default":
	default:
		panic(fmt.Sprintf("unknown stage: %s", stage))
	}

	if postInterpreter, ok := res.(basil.BlockPostInterpreter); ok {
		if err = postInterpreter.PostEval(ctx, stage); err != nil {
			return err
		}
	}

	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockValueRequiredInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockValueRequiredInterpreter) ValueParamName() basil.ID {
	return basil.ID("value")
}

func (i BlockValueRequiredInterpreter) BlockRegistry() parsley.NodeTransformerRegistry {
	var block basil.Block = &BlockValueRequired{}
	if b, ok := block.(basil.BlockRegistryAware); ok {
		return b.BlockRegistry()
	}

	return nil
}
