// Code generated by Basil. DO NOT EDIT.

package test

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// TestFunc1Interpreter is the basil interpreter for the testFunc1 function
type TestFunc1Interpreter struct {
	s schema.Schema
}

func (i TestFunc1Interpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name:   "str",
					Schema: &schema.String{},
				},
			},
			Result: &schema.String{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i TestFunc1Interpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	arg0, evalErr := parsley.EvaluateNode(ctx, arguments[0])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[0].Schema.ValidateValue(arg0); err != nil {
		return nil, parsley.NewError(arguments[0].Pos(), err)
	}
	var val0 = arg0.(string)

	return testFunc1(val0), nil
}
